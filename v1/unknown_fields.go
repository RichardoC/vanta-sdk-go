package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"reflect"
	"strings"
	"sync"
)

// UnknownFieldWarningf is called when typed JSON decoding encounters a JSON
// object field that is not represented by the SDK model. Set it to nil to
// suppress warnings.
var UnknownFieldWarningf = log.Printf

var (
	unknownFieldWarningsMu   sync.Mutex
	unknownFieldWarningsSeen = map[string]struct{}{}
)

var jsonRawMessageType = reflect.TypeFor[json.RawMessage]()

func decodeJSONBytes(data []byte, out any) error {
	if err := json.Unmarshal(data, out); err != nil {
		return err
	}

	warnUnknownJSONFields(data, out)
	return nil
}

func warnUnknownJSONFields(data []byte, out any) {
	if out == nil || UnknownFieldWarningf == nil {
		return
	}

	var raw any
	if err := json.Unmarshal(data, &raw); err != nil {
		return
	}

	t := reflect.TypeOf(out)
	if t == nil {
		return
	}
	rootName := typePathName(derefType(t))
	if rootName == "" {
		rootName = "response"
	}
	walkUnknownFields(derefType(t), raw, rootName)
}

func warnUnknownField(path string) {
	unknownFieldWarningsMu.Lock()
	if _, seen := unknownFieldWarningsSeen[path]; seen {
		unknownFieldWarningsMu.Unlock()
		return
	}
	unknownFieldWarningsSeen[path] = struct{}{}
	warnf := UnknownFieldWarningf
	unknownFieldWarningsMu.Unlock()

	if warnf != nil {
		warnf("vanta-sdk-go: unknown response field detected: %s", path)
	}
}

func walkUnknownFields(t reflect.Type, raw any, path string) {
	t = derefType(t)
	if t == nil || raw == nil {
		return
	}

	if t == jsonRawMessageType || isByteSliceType(t) {
		return
	}

	switch t.Kind() {
	case reflect.Struct:
		obj, ok := raw.(map[string]any)
		if !ok {
			return
		}
		fields := jsonFieldsForType(t)
		for key := range obj {
			if _, ok := fields[key]; !ok {
				warnUnknownField(path + "." + key)
			}
		}
		for key, fieldType := range fields {
			value, ok := obj[key]
			if !ok {
				continue
			}
			walkUnknownFields(fieldType, value, path+"."+key)
		}
	case reflect.Slice, reflect.Array:
		items, ok := raw.([]any)
		if !ok {
			return
		}
		for _, item := range items {
			walkUnknownFields(t.Elem(), item, path+"[]")
		}
	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			return
		}
		if derefType(t.Elem()).Kind() == reflect.Interface {
			return
		}
		obj, ok := raw.(map[string]any)
		if !ok {
			return
		}
		for key, value := range obj {
			walkUnknownFields(t.Elem(), value, path+"."+key)
		}
	}
}

func jsonFieldsForType(t reflect.Type) map[string]reflect.Type {
	fields := map[string]reflect.Type{}
	for field := range t.Fields() {
		if field.PkgPath != "" && !field.Anonymous {
			continue
		}
		if field.Anonymous {
			embedded := derefType(field.Type)
			if embedded == nil || embedded.Kind() != reflect.Struct {
				continue
			}
			maps.Copy(fields, jsonFieldsForType(embedded))
			continue
		}

		name := jsonFieldName(field)
		if name == "" {
			continue
		}
		fields[name] = field.Type
	}
	return fields
}

func jsonFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "-" {
		return ""
	}
	if tag == "" {
		return field.Name
	}
	name := strings.Split(tag, ",")[0]
	if name == "" {
		return field.Name
	}
	return name
}

func derefType(t reflect.Type) reflect.Type {
	for t != nil && t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func typePathName(t reflect.Type) string {
	if t == nil {
		return ""
	}
	if name := t.Name(); name != "" {
		return name
	}
	switch t.Kind() {
	case reflect.Struct:
		return "struct"
	case reflect.Slice, reflect.Array:
		return typePathName(t.Elem())
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", t.Key(), typePathName(t.Elem()))
	default:
		return t.String()
	}
}

func isByteSliceType(t reflect.Type) bool {
	return t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8
}

func resetUnknownFieldWarningsForTest() {
	unknownFieldWarningsMu.Lock()
	unknownFieldWarningsSeen = map[string]struct{}{}
	unknownFieldWarningsMu.Unlock()
}
