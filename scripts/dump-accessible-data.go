package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	vanta "github.com/richardoc/vanta-sdk-go/v1"
)

type callResult struct {
	Service string `json:"service"`
	Method  string `json:"method"`
	Status  string `json:"status"`
	File    string `json:"file,omitempty"`
	Pages   int    `json:"pages,omitempty"`
	Error   string `json:"error,omitempty"`
}

var safeName = regexp.MustCompile(`[^a-zA-Z0-9._-]+`)

func main() {
	ctx := context.Background()
	outDir := strings.TrimSpace(os.Getenv("VANTA_DUMP_DIR"))
	if outDir == "" {
		outDir = "vanta-api-dump"
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("create output dir: %v", err)
	}

	client, err := newVantaClient(ctx)
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	results, err := dumpAllCallableEndpoints(ctx, client, outDir)
	if err != nil {
		log.Fatalf("dump endpoints: %v", err)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Service == results[j].Service {
			return results[i].Method < results[j].Method
		}
		return results[i].Service < results[j].Service
	})

	if err := writeJSON(filepath.Join(outDir, "index.json"), results); err != nil {
		log.Fatalf("write index: %v", err)
	}

	success := 0
	errored := 0
	for _, r := range results {
		if r.Status == "ok" {
			success++
		} else {
			errored++
		}
	}

	fmt.Printf("Dump complete: %d success, %d errors\n", success, errored)
	fmt.Printf("Output directory: %s\n", outDir)
}

func newVantaClient(ctx context.Context) (*vanta.Client, error) {
	opts := make([]vanta.Option, 0, 2)
	if baseURL := strings.TrimSpace(os.Getenv("VANTA_BASE_URL")); baseURL != "" {
		opts = append(opts, vanta.WithBaseURL(baseURL))
	}

	if token := strings.TrimSpace(os.Getenv("VANTA_BEARER_TOKEN")); token != "" {
		opts = append(opts, vanta.WithTokenSource(vanta.StaticTokenSource(token)))
		return vanta.NewClient(opts...)
	}

	clientID := strings.TrimSpace(os.Getenv("VANTA_CLIENT_ID"))
	clientSecret := strings.TrimSpace(os.Getenv("VANTA_CLIENT_SECRET"))
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("set either VANTA_BEARER_TOKEN or both VANTA_CLIENT_ID and VANTA_CLIENT_SECRET")
	}

	scope := strings.TrimSpace(os.Getenv("VANTA_SCOPE"))
	if scope == "" {
		scope = "vanta-api.all:read"
	}

	ts, err := vanta.NewOAuthClientCredentialsTokenSource(vanta.OAuthClientCredentialsConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	})
	if err != nil {
		return nil, err
	}

	opts = append(opts, vanta.WithTokenSource(ts))
	return vanta.NewClient(opts...)
}

func dumpAllCallableEndpoints(ctx context.Context, client *vanta.Client, outDir string) ([]callResult, error) {
	servicesVal := reflect.ValueOf(client.Services).Elem()
	servicesType := servicesVal.Type()
	results := make([]callResult, 0, 256)

	for i := 0; i < servicesVal.NumField(); i++ {
		serviceField := servicesType.Field(i)
		serviceName := serviceField.Name
		serviceValue := servicesVal.Field(i)
		if serviceValue.IsNil() {
			continue
		}

		serviceMethods := listCallableMethods(serviceValue)
		for _, methodName := range serviceMethods {
			res := callServiceMethod(ctx, serviceName, serviceValue, methodName, outDir)
			results = append(results, res)
		}
	}

	return results, nil
}

func listCallableMethods(serviceValue reflect.Value) []string {
	serviceType := serviceValue.Type()
	includeMutations := strings.EqualFold(strings.TrimSpace(os.Getenv("VANTA_INCLUDE_MUTATIONS")), "1") ||
		strings.EqualFold(strings.TrimSpace(os.Getenv("VANTA_INCLUDE_MUTATIONS")), "true")

	methods := make([]string, 0, serviceType.NumMethod())
	for m := range serviceType.Methods() {
		m := m
		if !includeMutations && !isLikelyReadOnlyMethod(m.Name) {
			continue
		}
		// Use the bound method type (receiver omitted) for invocation validation.
		bound := serviceValue.MethodByName(m.Name)
		if !bound.IsValid() {
			continue
		}
		mt := bound.Type()
		if mt.NumIn() != 2 {
			continue
		}
		if mt.In(0) != reflect.TypeFor[context.Context]() {
			continue
		}
		if mt.In(1).Kind() != reflect.Pointer {
			continue
		}
		if mt.NumOut() != 2 {
			continue
		}
		if !mt.Out(1).Implements(reflect.TypeFor[error]()) {
			continue
		}
		methods = append(methods, m.Name)
	}
	sort.Strings(methods)
	return methods
}

func isLikelyReadOnlyMethod(name string) bool {
	return strings.HasPrefix(name, "Get") || strings.HasPrefix(name, "List")
}

func callServiceMethod(ctx context.Context, serviceName string, serviceValue reflect.Value, methodName, outDir string) callResult {
	if strings.HasPrefix(methodName, "List") {
		return callPaginatedMethod(ctx, serviceName, serviceValue, methodName, outDir)
	}

	result := callResult{Service: serviceName, Method: methodName, Status: "error"}
	method := serviceValue.MethodByName(methodName)
	if !method.IsValid() {
		result.Error = "method not found"
		return result
	}

	callCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
	defer cancel()

	var outs []reflect.Value
	func() {
		defer func() {
			if r := recover(); r != nil {
				result.Error = fmt.Sprintf("panic: %v", r)
			}
		}()
		outs = method.Call([]reflect.Value{reflect.ValueOf(callCtx), reflect.Zero(method.Type().In(1))})
	}()
	if result.Error != "" {
		return result
	}

	errVal := outs[1]
	if !errVal.IsNil() {
		result.Error = errVal.Interface().(error).Error()
		return result
	}

	respVal := outs[0]
	if !respVal.IsValid() || (respVal.Kind() == reflect.Pointer && respVal.IsNil()) {
		result.Status = "ok"
		result.File = ""
		return result
	}

	fileName := sanitize(fmt.Sprintf("%s_%s.json", serviceName, methodName))
	filePath := filepath.Join(outDir, fileName)

	if err := writeResponse(filePath, respVal.Interface()); err != nil {
		result.Error = fmt.Sprintf("write response: %v", err)
		return result
	}

	result.Status = "ok"
	result.File = fileName
	return result
}

func callPaginatedMethod(ctx context.Context, serviceName string, serviceValue reflect.Value, methodName, outDir string) callResult {
	result := callResult{Service: serviceName, Method: methodName, Status: "error"}
	method := serviceValue.MethodByName(methodName)
	if !method.IsValid() {
		result.Error = "method not found"
		return result
	}

	pageSize := 100
	if raw := strings.TrimSpace(os.Getenv("VANTA_PAGE_SIZE")); raw != "" {
		if n, err := fmt.Sscanf(raw, "%d", &pageSize); err != nil || n != 1 || pageSize < 1 {
			result.Error = fmt.Sprintf("invalid VANTA_PAGE_SIZE: %q", raw)
			return result
		}
		if pageSize > 100 {
			pageSize = 100
		}
	}

	cursor := ""
	pageNum := 0
	wroteAny := false

	for {
		pageNum++
		params := buildParams(method.Type().In(1), pageSize, cursor)

		callCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
		var outs []reflect.Value
		func() {
			defer func() {
				if r := recover(); r != nil {
					result.Error = fmt.Sprintf("panic: %v", r)
				}
			}()
			outs = method.Call([]reflect.Value{reflect.ValueOf(callCtx), params})
		}()
		cancel()
		if result.Error != "" {
			return result
		}

		errVal := outs[1]
		if !errVal.IsNil() {
			result.Error = errVal.Interface().(error).Error()
			return result
		}

		respVal := outs[0]
		if !respVal.IsValid() || (respVal.Kind() == reflect.Pointer && respVal.IsNil()) {
			result.Status = "ok"
			result.Pages = pageNum
			return result
		}

		fileName := sanitize(fmt.Sprintf("%s_%s_page_%03d.json", serviceName, methodName, pageNum))
		filePath := filepath.Join(outDir, fileName)
		if err := writeResponse(filePath, respVal.Interface()); err != nil {
			result.Error = fmt.Sprintf("write response: %v", err)
			return result
		}
		if !wroteAny {
			result.File = fileName
			wroteAny = true
		}

		hasNext, endCursor, ok := extractPageInfo(respVal.Interface())
		if !ok || !hasNext || endCursor == "" || endCursor == cursor {
			result.Status = "ok"
			result.Pages = pageNum
			return result
		}
		cursor = endCursor
	}
}

func buildParams(paramType reflect.Type, pageSize int, cursor string) reflect.Value {
	params := reflect.New(paramType.Elem())
	elem := params.Elem()

	if f := elem.FieldByName("PageSize"); f.IsValid() && f.CanSet() {
		if f.Kind() == reflect.Pointer && f.Type().Elem().Kind() == reflect.Int {
			v := reflect.New(f.Type().Elem())
			v.Elem().SetInt(int64(pageSize))
			f.Set(v)
		}
	}
	if cursor != "" {
		if f := elem.FieldByName("PageCursor"); f.IsValid() && f.CanSet() {
			if f.Kind() == reflect.Pointer && f.Type().Elem().Kind() == reflect.String {
				v := reflect.New(f.Type().Elem())
				v.Elem().SetString(cursor)
				f.Set(v)
			}
		}
	}
	return params
}

func extractPageInfo(payload any) (bool, string, bool) {
	var b []byte
	switch v := payload.(type) {
	case json.RawMessage:
		b = v
	default:
		var err error
		b, err = json.Marshal(v)
		if err != nil {
			return false, "", false
		}
	}

	var root map[string]any
	if err := json.Unmarshal(b, &root); err != nil {
		return false, "", false
	}
	results, ok := root["results"].(map[string]any)
	if !ok {
		return false, "", false
	}
	pageInfo, ok := results["pageInfo"].(map[string]any)
	if !ok {
		return false, "", false
	}
	hasNext, _ := pageInfo["hasNextPage"].(bool)
	endCursor, _ := pageInfo["endCursor"].(string)
	return hasNext, endCursor, true
}

func writeResponse(path string, payload any) error {
	switch v := payload.(type) {
	case json.RawMessage:
		if len(v) == 0 {
			return os.WriteFile(path, []byte("null\n"), 0o644)
		}
		var pretty any
		if err := json.Unmarshal(v, &pretty); err == nil {
			return writeJSON(path, pretty)
		}
		return os.WriteFile(path, append(v, '\n'), 0o644)
	default:
		return writeJSON(path, v)
	}
}

func writeJSON(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	b = append(b, '\n')
	return os.WriteFile(path, b, 0o644)
}

func sanitize(s string) string {
	clean := safeName.ReplaceAllString(s, "_")
	clean = strings.Trim(clean, "._-")
	if clean == "" {
		return "response.json"
	}
	return clean
}
