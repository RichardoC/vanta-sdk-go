package v1

import "testing"

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
