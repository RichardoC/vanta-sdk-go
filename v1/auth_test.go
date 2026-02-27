package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestStaticTokenSource(t *testing.T) {
	tok, err := StaticTokenSource("abc").Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.AccessToken != "abc" {
		t.Fatalf("access token = %q, want abc", tok.AccessToken)
	}
}

func TestOAuthClientCredentialsTokenSourceCachesToken(t *testing.T) {
	var calls int32
	mockClient := &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			atomic.AddInt32(&calls, 1)
			if r.Method != http.MethodPost {
				t.Fatalf("method = %s, want POST", r.Method)
			}
			payload, _ := json.Marshal(map[string]any{
				"access_token": "token-1",
				"token_type":   "Bearer",
				"expires_in":   3600,
			})
			return &http.Response{
				StatusCode: http.StatusOK,
				Status:     "200 OK",
				Header:     http.Header{"Content-Type": []string{"application/json"}},
				Body:       io.NopCloser(strings.NewReader(string(payload))),
			}, nil
		}),
	}

	src, err := NewOAuthClientCredentialsTokenSource(OAuthClientCredentialsConfig{
		ClientID:     "id",
		ClientSecret: "secret",
		AuthURL:      "https://api.vanta.com/oauth/token",
		HTTPClient:   mockClient,
		RefreshSkew:  time.Minute,
	})
	if err != nil {
		t.Fatalf("unexpected constructor error: %v", err)
	}

	ctx := context.Background()
	tok1, err := src.Token(ctx)
	if err != nil {
		t.Fatalf("first token error: %v", err)
	}
	tok2, err := src.Token(ctx)
	if err != nil {
		t.Fatalf("second token error: %v", err)
	}

	if tok1.AccessToken != "token-1" || tok2.AccessToken != "token-1" {
		t.Fatalf("unexpected tokens: tok1=%q tok2=%q", tok1.AccessToken, tok2.AccessToken)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("oauth endpoint calls = %d, want 1", got)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
