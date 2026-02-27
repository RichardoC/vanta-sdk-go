package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Token represents an OAuth access token.
type Token struct {
	AccessToken string
	TokenType   string
	Expiry      time.Time
}

// Valid reports whether the token appears usable for requests.
func (t Token) Valid() bool {
	if t.AccessToken == "" {
		return false
	}
	if t.Expiry.IsZero() {
		return true
	}
	return time.Now().Before(t.Expiry)
}

// TokenSource resolves bearer tokens for API requests.
type TokenSource interface {
	Token(ctx context.Context) (Token, error)
}

// StaticTokenSource returns a fixed token.
type StaticTokenSource string

// Token returns the static token value.
func (s StaticTokenSource) Token(context.Context) (Token, error) {
	if s == "" {
		return Token{}, fmt.Errorf("static token is empty")
	}
	return Token{AccessToken: string(s), TokenType: "Bearer"}, nil
}

// OAuthClientCredentialsConfig configures Vanta OAuth client credentials flow.
type OAuthClientCredentialsConfig struct {
	ClientID     string
	ClientSecret string
	Scope        string
	AuthURL      string
	HTTPClient   *http.Client
	RefreshSkew  time.Duration
}

// OAuthClientCredentialsTokenSource fetches and caches OAuth tokens.
//
// Vanta only allows one active token per API app credentials. This source uses
// synchronized refresh to avoid concurrent token churn from the same process.
type OAuthClientCredentialsTokenSource struct {
	httpClient   *http.Client
	authURL      string
	clientID     string
	clientSecret string
	scope        string
	refreshSkew  time.Duration

	mu    sync.Mutex
	token Token
}

// NewOAuthClientCredentialsTokenSource builds a token source backed by OAuth
// client_credentials grant.
func NewOAuthClientCredentialsTokenSource(cfg OAuthClientCredentialsConfig) (*OAuthClientCredentialsTokenSource, error) {
	if cfg.ClientID == "" {
		return nil, fmt.Errorf("client ID must not be empty")
	}
	if cfg.ClientSecret == "" {
		return nil, fmt.Errorf("client secret must not be empty")
	}
	if cfg.AuthURL == "" {
		cfg.AuthURL = defaultAPIBaseURL + defaultOAuthPath
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 15 * time.Second}
	}
	if cfg.RefreshSkew <= 0 {
		cfg.RefreshSkew = 60 * time.Second
	}

	return &OAuthClientCredentialsTokenSource{
		httpClient:   cfg.HTTPClient,
		authURL:      cfg.AuthURL,
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		scope:        cfg.Scope,
		refreshSkew:  cfg.RefreshSkew,
	}, nil
}

// Token returns a valid cached token or refreshes it when required.
func (s *OAuthClientCredentialsTokenSource) Token(ctx context.Context) (Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.token.AccessToken != "" {
		expiresSoon := !s.token.Expiry.IsZero() && time.Until(s.token.Expiry) <= s.refreshSkew
		if !expiresSoon && s.token.Valid() {
			return s.token, nil
		}
	}

	tok, err := s.fetchToken(ctx)
	if err != nil {
		return Token{}, err
	}
	s.token = tok
	return tok, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (s *OAuthClientCredentialsTokenSource) fetchToken(ctx context.Context) (Token, error) {
	payload := map[string]string{
		"client_id":     s.clientID,
		"client_secret": s.clientSecret,
		"grant_type":    "client_credentials",
	}
	if s.scope != "" {
		payload["scope"] = s.scope
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return Token{}, fmt.Errorf("marshal oauth payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.authURL, bytes.NewReader(body))
	if err != nil {
		return Token{}, fmt.Errorf("build oauth request: %w", err)
	}
	req.Header.Set("Content-Type", contentTypeJSON)
	req.Header.Set("Accept", contentTypeJSON)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return Token{}, fmt.Errorf("oauth request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return Token{}, decodeAPIError(resp)
	}

	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return Token{}, fmt.Errorf("decode oauth token response: %w", err)
	}
	if tr.AccessToken == "" {
		return Token{}, fmt.Errorf("oauth token response missing access_token")
	}

	tok := Token{AccessToken: tr.AccessToken, TokenType: tr.TokenType}
	if tr.ExpiresIn > 0 {
		tok.Expiry = time.Now().Add(time.Duration(tr.ExpiresIn) * time.Second)
	}
	if tok.TokenType == "" {
		tok.TokenType = "Bearer"
	}
	return tok, nil
}
