package v1

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestJoinURLPath(t *testing.T) {
	tests := []struct {
		name string
		base string
		rel  string
		want string
	}{
		{name: "empty base", base: "", rel: "v1/controls", want: "/v1/controls"},
		{name: "trim slash", base: "/api/", rel: "/v1/controls", want: "/api/v1/controls"},
		{name: "empty relative", base: "/api", rel: "", want: "/api"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := joinURLPath(tt.base, tt.rel)
			if got != tt.want {
				t.Fatalf("joinURLPath(%q, %q) = %q, want %q", tt.base, tt.rel, got, tt.want)
			}
		})
	}
}

func TestDefaultBaseURLTargetsV1PeopleEndpoint(t *testing.T) {
	var gotURL string
	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			gotURL = r.URL.String()
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(
					`{"results":{"data":[],"pageInfo":{"hasNextPage":false,"hasPreviousPage":false,"startCursor":"","endCursor":""}}}`,
				)),
			}, nil
		}),
	}

	c, err := NewClient(WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	_, err = c.Services.People.ListPeople(context.Background(), &PeopleListPeopleParams{})
	if err != nil {
		t.Fatalf("ListPeople returned error: %v", err)
	}

	want := "https://api.vanta.com/v1/people"
	if gotURL != want {
		t.Fatalf("ListPeople requested %q, want %q", gotURL, want)
	}
}

func TestDoJSONWarnsOnUnknownFields(t *testing.T) {
	resetUnknownFieldWarningsForTest()
	t.Cleanup(resetUnknownFieldWarningsForTest)

	var warnings []string
	previous := UnknownFieldWarningf
	UnknownFieldWarningf = func(format string, args ...any) {
		warnings = append(warnings, fmt.Sprintf(format, args...))
	}
	t.Cleanup(func() {
		UnknownFieldWarningf = previous
	})

	httpClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body: io.NopCloser(strings.NewReader(
					`{"item":{"name":"ok","newNestedField":true},"newRootField":1}`,
				)),
			}, nil
		}),
	}

	c, err := NewClient(WithHTTPClient(httpClient))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	req, err := c.newRequest(context.Background(), http.MethodGet, "/test", url.Values{}, nil)
	if err != nil {
		t.Fatalf("newRequest returned error: %v", err)
	}

	type nested struct {
		Name string `json:"name"`
	}
	type sampleResponse struct {
		Item nested `json:"item"`
	}

	var out sampleResponse
	if err := c.doJSON(req, &out); err != nil {
		t.Fatalf("doJSON returned error: %v", err)
	}
	if out.Item.Name != "ok" {
		t.Fatalf("decoded response = %+v", out)
	}

	joined := strings.Join(warnings, "\n")
	for _, want := range []string{
		"sampleResponse.newRootField",
		"sampleResponse.item.newNestedField",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("warnings %q do not contain %q", joined, want)
		}
	}
}
