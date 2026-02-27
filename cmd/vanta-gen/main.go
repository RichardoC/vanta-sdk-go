package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type postmanCollection struct {
	Item []postmanItem `json:"item"`
}

type postmanItem struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Item        []postmanItem `json:"item"`
	Request     *postmanReq   `json:"request"`
	Response    []postmanResp `json:"response"`
}

type postmanReq struct {
	Method      string       `json:"method"`
	Description string       `json:"description"`
	URL         postmanURL   `json:"url"`
	Body        *postmanBody `json:"body"`
}

type postmanBody struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
}

type postmanResp struct {
	Body string `json:"body"`
}

type postmanURL struct {
	Raw   string             `json:"raw"`
	Path  []string           `json:"path"`
	Query []postmanQueryItem `json:"query"`
}

type postmanQueryItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type operation struct {
	ServiceName  string
	MethodName   string
	Method       string
	Path         string
	Description  string
	QueryParams  []queryParam
	PathParams   []string
	HasJSONBody  bool
	HasMultipart bool
	ReqShape     *jsonShape
	RespShape    *jsonShape
}

type queryParam struct {
	Name   string
	GoName string
	Type   string
}

var identSanitizer = regexp.MustCompile(`[^a-zA-Z0-9]+`)
var camelBoundary = regexp.MustCompile(`([a-z0-9])([A-Z])`)
var stopwordSet = map[string]bool{"a": true, "an": true, "the": true}

type jsonField struct {
	Name  string
	Type  string
	Tag   string
	Order int
}

type jsonShape struct {
	IsObject bool
	Fields   []jsonField
}

func main() {
	collectionPath := filepath.Join("Vanta Postman Env & Collection", "Vanta API.postman_collection.json")
	b, err := os.ReadFile(collectionPath)
	if err != nil {
		fatalf("read collection: %v", err)
	}

	var c postmanCollection
	if err := json.Unmarshal(b, &c); err != nil {
		fatalf("parse collection: %v", err)
	}

	var ops []operation
	walkItems(c.Item, nil, &ops)
	ops = dedupeOperations(ops)
	if len(ops) == 0 {
		fatalf("no operations found")
	}

	services := map[string][]operation{}
	for _, op := range ops {
		services[op.ServiceName] = append(services[op.ServiceName], op)
	}

	var serviceNames []string
	for name := range services {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)
	for k := range services {
		sort.SliceStable(services[k], func(i, j int) bool {
			if services[k][i].MethodName == services[k][j].MethodName {
				return services[k][i].Path < services[k][j].Path
			}
			return services[k][i].MethodName < services[k][j].MethodName
		})
		services[k] = uniqueMethodNames(services[k])
	}

	generated, err := render(services, serviceNames)
	if err != nil {
		fatalf("render: %v", err)
	}

	outPath := filepath.Join("v1", "generated_services.go")
	if err := os.WriteFile(outPath, generated, 0o644); err != nil {
		fatalf("write generated file: %v", err)
	}

	fmt.Printf("generated %s with %d operations across %d services\n", outPath, len(ops), len(serviceNames))
}

func walkItems(items []postmanItem, parents []string, out *[]operation) {
	for _, it := range items {
		trail := append(parents[:len(parents):len(parents)], it.Name)
		if it.Request != nil {
			op, ok := buildOperation(trail, it)
			if ok {
				*out = append(*out, op)
			}
		}
		if len(it.Item) > 0 {
			walkItems(it.Item, trail, out)
		}
	}
}

func buildOperation(trail []string, it postmanItem) (operation, bool) {
	rawPath := it.Request.URL.Raw
	if rawPath == "" {
		return operation{}, false
	}

	normalizedPath, rawQuery := normalizeRawURL(rawPath)
	if normalizedPath == "" {
		return operation{}, false
	}

	serviceName := serviceFromPath(normalizedPath)
	methodName := goExportedName(it.Name, true)
	if methodName == "" {
		methodName = goExportedName(strings.Join(trail, " "), true)
	}
	if methodName == "" {
		methodName = strings.Title(strings.ToLower(it.Request.Method))
	}
	if methodName == "List" {
		methodName = "ListItems"
	}

	queryItems := it.Request.URL.Query
	if len(queryItems) == 0 && rawQuery != "" {
		queryItems = parseQueryFromRaw(rawQuery)
	}

	seenQuery := map[string]bool{}
	qps := make([]queryParam, 0, len(queryItems))
	for _, q := range queryItems {
		if q.Key == "" || seenQuery[q.Key] {
			continue
		}
		seenQuery[q.Key] = true
		qps = append(qps, queryParam{
			Name:   q.Key,
			GoName: goExportedName(q.Key, false),
			Type:   inferQueryType(q.Value),
		})
	}

	desc := strings.TrimSpace(it.Request.Description)
	if desc == "" {
		desc = strings.TrimSpace(it.Description)
	}
	pathParams := extractPathParams(normalizedPath)
	bodyMode := ""
	if it.Request.Body != nil {
		bodyMode = it.Request.Body.Mode
	}
	reqShape := shapeFromRawJSON("")
	if bodyMode == "raw" && it.Request.Body != nil {
		reqShape = shapeFromRawJSON(it.Request.Body.Raw)
	}

	respShape := shapeFromRawJSON("")
	for _, rsp := range it.Response {
		respShape = shapeFromRawJSON(rsp.Body)
		if respShape != nil {
			break
		}
	}

	return operation{
		ServiceName:  serviceName,
		MethodName:   methodName,
		Method:       strings.ToUpper(it.Request.Method),
		Path:         normalizedPath,
		Description:  oneLine(desc),
		QueryParams:  qps,
		PathParams:   pathParams,
		HasJSONBody:  bodyMode == "raw",
		HasMultipart: bodyMode == "formdata",
		ReqShape:     reqShape,
		RespShape:    respShape,
	}, true
}

func normalizeRawURL(raw string) (path, rawQuery string) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "{{baseUrl}}")
	raw = strings.TrimPrefix(raw, "{{authUrl}}")
	if raw == "" {
		return "", ""
	}
	if !strings.HasPrefix(raw, "/") {
		raw = "/" + raw
	}
	parts := strings.SplitN(raw, "?", 2)
	path = parts[0]
	if len(parts) > 1 {
		rawQuery = parts[1]
	}
	return path, rawQuery
}

func parseQueryFromRaw(rawQuery string) []postmanQueryItem {
	parts := strings.Split(rawQuery, "&")
	items := make([]postmanQueryItem, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 1 {
			items = append(items, postmanQueryItem{Key: kv[0]})
			continue
		}
		items = append(items, postmanQueryItem{Key: kv[0], Value: kv[1]})
	}
	return items
}

func serviceFromPath(path string) string {
	trimmed := strings.TrimPrefix(path, "/")
	parts := strings.Split(trimmed, "/")
	if len(parts) == 0 || parts[0] == "" {
		return "Root"
	}
	if parts[0] == "v1" && len(parts) > 1 {
		return goExportedName(parts[1], false)
	}
	return goExportedName(parts[0], false)
}

func inferQueryType(sample string) string {
	sample = strings.TrimSpace(sample)
	if sample == "" {
		return "string"
	}
	if sample == "true" || sample == "false" {
		return "bool"
	}
	if _, err := strconv.Atoi(sample); err == nil {
		return "int"
	}
	return "string"
}

func extractPathParams(path string) []string {
	parts := strings.Split(path, "/")
	params := []string{}
	for _, p := range parts {
		if strings.HasPrefix(p, ":") && len(p) > 1 {
			params = append(params, p[1:])
		}
	}
	return params
}

func goExportedName(s string, dropStopwords bool) string {
	s = strings.ReplaceAll(s, "'s", "s")
	s = strings.ReplaceAll(s, "'", "")
	s = camelBoundary.ReplaceAllString(s, `$1 $2`)
	s = identSanitizer.ReplaceAllString(s, " ")
	parts := strings.Fields(strings.TrimSpace(strings.ToLower(s)))
	if len(parts) == 0 {
		return ""
	}
	cleaned := make([]string, 0, len(parts))
	for _, part := range parts {
		if dropStopwords && stopwordSet[part] {
			continue
		}
		cleaned = append(cleaned, strings.Title(part))
	}
	if len(cleaned) == 0 {
		cleaned = parts
		for i := range cleaned {
			cleaned[i] = strings.Title(cleaned[i])
		}
	}
	name := strings.Join(cleaned, "")
	replacer := strings.NewReplacer("Id", "ID", "Url", "URL", "Api", "API", "Oauth", "OAuth")
	name = replacer.Replace(name)
	if name == "" {
		return ""
	}
	if name[0] >= '0' && name[0] <= '9' {
		name = "N" + name
	}
	return name
}

func oneLine(s string) string {
	if s == "" {
		return s
	}
	return strings.Join(strings.Fields(s), " ")
}

func dedupeOperations(ops []operation) []operation {
	seen := map[string]operation{}
	order := make([]string, 0, len(ops))
	for _, op := range ops {
		key := operationSignature(op)
		if existing, ok := seen[key]; ok {
			if existing.Description == "" && op.Description != "" {
				existing.Description = op.Description
			}
			if existing.ReqShape == nil && op.ReqShape != nil {
				existing.ReqShape = op.ReqShape
			}
			if existing.RespShape == nil && op.RespShape != nil {
				existing.RespShape = op.RespShape
			}
			seen[key] = existing
			continue
		}
		seen[key] = op
		order = append(order, key)
	}
	out := make([]operation, 0, len(order))
	for _, key := range order {
		out = append(out, seen[key])
	}
	return out
}

func operationSignature(op operation) string {
	var qNames []string
	for _, q := range op.QueryParams {
		qNames = append(qNames, q.Name+":"+q.Type)
	}
	sort.Strings(qNames)

	req := "none"
	if op.ReqShape != nil {
		req = shapeSignature(op.ReqShape)
	}
	resp := "none"
	if op.RespShape != nil {
		resp = shapeSignature(op.RespShape)
	}

	return strings.Join([]string{
		op.Method,
		op.Path,
		strings.Join(qNames, ","),
		fmt.Sprintf("json:%t", op.HasJSONBody),
		fmt.Sprintf("multipart:%t", op.HasMultipart),
		"req:" + req,
		"resp:" + resp,
	}, "|")
}

func shapeSignature(s *jsonShape) string {
	if s == nil || !s.IsObject {
		return "none"
	}
	parts := make([]string, 0, len(s.Fields))
	for _, f := range s.Fields {
		parts = append(parts, f.Tag+":"+f.Type)
	}
	sort.Strings(parts)
	return strings.Join(parts, ",")
}

func shapeFromRawJSON(raw string) *jsonShape {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	var v any
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		return nil
	}
	obj, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	fields := make([]jsonField, 0, len(obj))
	i := 0
	for k, val := range obj {
		fields = append(fields, jsonField{
			Name:  goExportedName(k, false),
			Type:  inferJSONValueType(val),
			Tag:   k,
			Order: i,
		})
		i++
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })
	return &jsonShape{IsObject: true, Fields: fields}
}

func inferJSONValueType(v any) string {
	switch t := v.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	case []any:
		if len(t) == 0 {
			return "[]any"
		}
		return "[]" + inferJSONValueType(t[0])
	case map[string]any:
		return "map[string]any"
	default:
		return "any"
	}
}

func zeroReturnExpr(typed bool) string {
	if typed {
		return "nil"
	}
	return "nil"
}

func uniqueMethodNames(ops []operation) []operation {
	used := map[string]int{}
	emitted := map[string]bool{}
	for i := range ops {
		name := ops[i].MethodName
		if used[name] == 0 {
			used[name] = 1
			emitted[name] = true
			continue
		}
		op := &ops[i]
		suffix := disambiguatorSuffix(op.Path, op.ServiceName)
		candidate := name + suffix
		if candidate == name {
			candidate = name + "Alt"
		}
		if emitted[candidate] {
			n := 2
			for emitted[fmt.Sprintf("%s%d", candidate, n)] {
				n++
			}
			candidate = fmt.Sprintf("%s%d", candidate, n)
		}
		op.MethodName = candidate
		used[name]++
		emitted[candidate] = true
	}
	return ops
}

func disambiguatorSuffix(path, serviceName string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	candidates := make([]string, 0, len(parts))
	serviceLower := strings.ToLower(serviceName)
	for _, p := range parts {
		if p == "" || p == "v1" || strings.HasPrefix(p, ":") {
			continue
		}
		if strings.EqualFold(p, serviceLower) {
			continue
		}
		candidates = append(candidates, p)
	}
	if len(candidates) == 0 {
		return ""
	}
	if len(candidates) > 2 {
		candidates = candidates[len(candidates)-2:]
	}
	return "For" + goExportedName(strings.Join(candidates, " "), true)
}

func render(services map[string][]operation, serviceNames []string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("// Code generated by cmd/vanta-gen. DO NOT EDIT.\n")
	b.WriteString("\n")
	b.WriteString("package v1\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"encoding/json\"\n")
	b.WriteString("\t\"fmt\"\n")
	b.WriteString("\t\"net/url\"\n")
	b.WriteString("\t\"strings\"\n")
	b.WriteString(")\n\n")

	b.WriteString("// Services is the generated service registry for Vanta APIs.\n")
	b.WriteString("type Services struct {\n")
	for _, svc := range serviceNames {
		fmt.Fprintf(&b, "\t%s *%sService\n", svc, svc)
	}
	b.WriteString("}\n\n")

	b.WriteString("func newGeneratedServices(c *Client) *Services {\n")
	b.WriteString("\tservices := &Services{}\n")
	for _, svc := range serviceNames {
		fmt.Fprintf(&b, "\tservices.%s = &%sService{client: c}\n", svc, svc)
	}
	b.WriteString("\treturn services\n")
	b.WriteString("}\n\n")

	for _, svc := range serviceNames {
		ops := services[svc]
		fmt.Fprintf(&b, "// %sService groups %d endpoint methods under the %q API segment.\n", svc, len(ops), svc)
		fmt.Fprintf(&b, "type %sService struct {\n\tclient *Client\n}\n\n", svc)

		for _, op := range ops {
			paramStruct := svc + op.MethodName + "Params"
			reqBodyType := svc + op.MethodName + "RequestBody"
			respType := svc + op.MethodName + "Response"
			hasTypedResp := op.RespShape != nil && op.RespShape.IsObject

			if op.HasJSONBody && op.ReqShape != nil && op.ReqShape.IsObject {
				fmt.Fprintf(&b, "type %s struct {\n", reqBodyType)
				for _, f := range op.ReqShape.Fields {
					fmt.Fprintf(&b, "\t%s %s `json:\"%s\"`\n", f.Name, f.Type, f.Tag)
				}
				b.WriteString("}\n\n")
			}
			if hasTypedResp {
				fmt.Fprintf(&b, "type %s struct {\n", respType)
				for _, f := range op.RespShape.Fields {
					fmt.Fprintf(&b, "\t%s %s `json:\"%s\"`\n", f.Name, f.Type, f.Tag)
				}
				b.WriteString("}\n\n")
			}

			if len(op.PathParams) > 0 || len(op.QueryParams) > 0 || op.HasJSONBody || op.HasMultipart {
				fmt.Fprintf(&b, "type %s struct {\n", paramStruct)
				for _, p := range op.PathParams {
					fmt.Fprintf(&b, "\t%s string\n", goExportedName(p, false))
				}
				for _, q := range op.QueryParams {
					fmt.Fprintf(&b, "\t%s *%s\n", q.GoName, q.Type)
				}
				if op.HasJSONBody {
					if op.ReqShape != nil && op.ReqShape.IsObject {
						fmt.Fprintf(&b, "\tBody *%s\n", reqBodyType)
					} else {
						b.WriteString("\tBody json.RawMessage\n")
					}
				}
				if op.HasMultipart {
					b.WriteString("\t// FormData maps multipart field names to values.\n")
					b.WriteString("\tFormData map[string]string\n")
				}
				b.WriteString("}\n\n")
			}

			doc := op.Description
			if doc == "" {
				doc = fmt.Sprintf("%s performs %s %s.", op.MethodName, op.Method, op.Path)
			}
			fmt.Fprintf(&b, "// %s %s\n", op.MethodName, doc)

			returnType := "json.RawMessage"
			if hasTypedResp {
				returnType = "*" + respType
			}
			if len(op.PathParams) > 0 || len(op.QueryParams) > 0 || op.HasJSONBody || op.HasMultipart {
				fmt.Fprintf(&b, "func (s *%sService) %s(ctx context.Context, params *%s) (%s, error) {\n", svc, op.MethodName, paramStruct, returnType)
				b.WriteString("\tif params == nil {\n\t\tparams = &" + paramStruct + "{}\n\t}\n")
			} else {
				fmt.Fprintf(&b, "func (s *%sService) %s(ctx context.Context) (%s, error) {\n", svc, op.MethodName, returnType)
			}

			fmt.Fprintf(&b, "\tpath := %q\n", op.Path)
			for _, p := range op.PathParams {
				goName := goExportedName(p, false)
				fmt.Fprintf(&b, "\tif params.%s == \"\" {\n\t\treturn %s, fmt.Errorf(%q)\n\t}\n", goName, zeroReturnExpr(hasTypedResp), p+" is required")
				fmt.Fprintf(&b, "\tpath = strings.ReplaceAll(path, %q, url.PathEscape(params.%s))\n", ":"+p, goName)
			}

			b.WriteString("\tquery := url.Values{}\n")
			for _, q := range op.QueryParams {
				fmt.Fprintf(&b, "\tif params.%s != nil {\n\t\tquery.Set(%q, fmt.Sprint(*params.%s))\n\t}\n", q.GoName, q.Name, q.GoName)
			}

			if op.HasMultipart {
				b.WriteString("\treq, err := s.client.newMultipartRequest(ctx, \"" + op.Method + "\", path, query, params.FormData)\n")
				b.WriteString("\tif err != nil {\n\t\treturn " + zeroReturnExpr(hasTypedResp) + ", err\n\t}\n")
				if hasTypedResp {
					fmt.Fprintf(&b, "\tout := &%s{}\n", respType)
					b.WriteString("\tif err := s.client.doJSON(req, out); err != nil {\n\t\treturn nil, err\n\t}\n")
				} else {
					b.WriteString("\tvar out json.RawMessage\n")
					b.WriteString("\tif err := s.client.doJSON(req, &out); err != nil {\n\t\treturn nil, err\n\t}\n")
				}
				b.WriteString("\treturn out, nil\n")
				b.WriteString("}\n\n")
				continue
			}

			if len(op.PathParams) > 0 || len(op.QueryParams) > 0 || op.HasJSONBody {
				if op.HasJSONBody {
					b.WriteString("\treq, err := s.client.newRequest(ctx, \"" + op.Method + "\", path, query, params.Body)\n")
				} else {
					b.WriteString("\treq, err := s.client.newRequest(ctx, \"" + op.Method + "\", path, query, nil)\n")
				}
			} else {
				b.WriteString("\treq, err := s.client.newRequest(ctx, \"" + op.Method + "\", path, query, nil)\n")
			}
			b.WriteString("\tif err != nil {\n\t\treturn " + zeroReturnExpr(hasTypedResp) + ", err\n\t}\n")
			if hasTypedResp {
				fmt.Fprintf(&b, "\tout := &%s{}\n", respType)
				b.WriteString("\tif err := s.client.doJSON(req, out); err != nil {\n\t\treturn nil, err\n\t}\n")
			} else {
				b.WriteString("\tvar out json.RawMessage\n")
				b.WriteString("\tif err := s.client.doJSON(req, &out); err != nil {\n\t\treturn nil, err\n\t}\n")
			}
			b.WriteString("\treturn out, nil\n")
			b.WriteString("}\n\n")
		}
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format generated source: %w", err)
	}
	return formatted, nil
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
