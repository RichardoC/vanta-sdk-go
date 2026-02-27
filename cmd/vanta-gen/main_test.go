package main

import (
	"strings"
	"testing"
)

func TestGoExportedNameStopwordsAndInitialisms(t *testing.T) {
	tests := []struct {
		in   string
		drop bool
		out  string
	}{
		{in: "Get control by an ID", drop: true, out: "GetControlByID"},
		{in: "List a control's documents", drop: true, out: "ListControlsDocuments"},
		{in: "oauth token url", drop: false, out: "OAuthTokenURL"},
		{in: "controlId", drop: false, out: "ControlID"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := goExportedName(tt.in, tt.drop)
			if got != tt.out {
				t.Fatalf("goExportedName(%q, %v) = %q, want %q", tt.in, tt.drop, got, tt.out)
			}
		})
	}
}

func TestDedupeOperationsByMethodAndPath(t *testing.T) {
	ops := []operation{
		{ServiceName: "OAuth", MethodName: "CreateToken", Method: "POST", Path: "/oauth/token", Description: "first"},
		{ServiceName: "OAuth", MethodName: "CreateTokenFromManage", Method: "POST", Path: "/oauth/token", Description: "second"},
		{ServiceName: "Controls", MethodName: "ListControls", Method: "GET", Path: "/controls"},
	}

	deduped := dedupeOperations(ops)
	if len(deduped) != 2 {
		t.Fatalf("len(deduped) = %d, want 2", len(deduped))
	}
	if deduped[0].Method != "POST" || deduped[0].Path != "/oauth/token" {
		t.Fatalf("unexpected first deduped op: %+v", deduped[0])
	}
	if deduped[1].Method != "GET" || deduped[1].Path != "/controls" {
		t.Fatalf("unexpected second deduped op: %+v", deduped[1])
	}
}

func TestDedupeOperationsPreservesDifferentSignatures(t *testing.T) {
	ops := []operation{
		{
			ServiceName: "Foo",
			MethodName:  "Update",
			Method:      "PATCH",
			Path:        "/foo/:id",
			HasJSONBody: true,
			ReqShape: &jsonShape{
				IsObject: true,
				Fields:   []jsonField{{Name: "Name", Type: "string", Tag: "name"}},
			},
		},
		{
			ServiceName: "Foo",
			MethodName:  "Update",
			Method:      "PATCH",
			Path:        "/foo/:id",
			HasJSONBody: true,
			ReqShape: &jsonShape{
				IsObject: true,
				Fields:   []jsonField{{Name: "Owner", Type: "string", Tag: "owner"}},
			},
		},
	}
	deduped := dedupeOperations(ops)
	if len(deduped) != 2 {
		t.Fatalf("len(deduped) = %d, want 2", len(deduped))
	}
}

func TestUniqueMethodNamesDisambiguatesByPath(t *testing.T) {
	ops := []operation{
		{ServiceName: "Foo", MethodName: "Update", Method: "PATCH", Path: "/foo/:id"},
		{ServiceName: "Foo", MethodName: "Update", Method: "PATCH", Path: "/foo/:id/state"},
	}
	out := uniqueMethodNames(ops)
	if out[1].MethodName != "UpdateForState" {
		t.Fatalf("unexpected disambiguated name: %q", out[1].MethodName)
	}
}

func TestRenderNamingGoldenSlice(t *testing.T) {
	services := map[string][]operation{
		"Controls": {
			{
				ServiceName: "Controls",
				MethodName:  goExportedName("Get control by an ID", true),
				Method:      "GET",
				Path:        "/controls/:controlId",
				Description: "Get a control by an ID.",
				PathParams:  []string{"controlId"},
				RespShape: &jsonShape{
					IsObject: true,
					Fields:   []jsonField{{Name: "ID", Type: "string", Tag: "id"}},
				},
			},
			{
				ServiceName: "Controls",
				MethodName:  goExportedName("List a control's documents", true),
				Method:      "GET",
				Path:        "/controls/:controlId/documents",
				Description: "List a control's documents.",
				PathParams:  []string{"controlId"},
				RespShape: &jsonShape{
					IsObject: true,
					Fields:   []jsonField{{Name: "Results", Type: "map[string]any", Tag: "results"}},
				},
			},
		},
	}
	out, err := render(services, []string{"Controls"})
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}
	got := string(out)
	mustContain := []string{
		"func (s *ControlsService) GetControlByID(",
		"func (s *ControlsService) ListControlsDocuments(",
		"ControlID string",
	}
	for _, want := range mustContain {
		if !strings.Contains(got, want) {
			t.Fatalf("render output missing %q", want)
		}
	}
}
