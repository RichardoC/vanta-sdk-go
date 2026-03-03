package v1

import (
	"context"
	"io"
	"net/http"
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
