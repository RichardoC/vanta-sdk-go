package v1

import (
	"errors"
	"net/http"
	"time"
)

type config struct {
	httpClient  *http.Client
	baseURL     string
	tokenSource TokenSource
	userAgent   string
}

func defaultConfig() *config {
	return &config{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    defaultAPIBaseURL,
		userAgent:  defaultUserAgent,
	}
}

// Option configures a client.
type Option interface {
	apply(*config) error
}

type optionFunc func(*config) error

func (f optionFunc) apply(cfg *config) error { return f(cfg) }

// WithBaseURL overrides the default API base URL.
func WithBaseURL(baseURL string) Option {
	return optionFunc(func(cfg *config) error {
		if baseURL == "" {
			return errors.New("base URL must not be empty")
		}
		cfg.baseURL = baseURL
		return nil
	})
}

// WithHTTPClient provides a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return optionFunc(func(cfg *config) error {
		if client == nil {
			return errors.New("http client must not be nil")
		}
		cfg.httpClient = client
		return nil
	})
}

// WithTokenSource sets the token source used for bearer auth.
func WithTokenSource(source TokenSource) Option {
	return optionFunc(func(cfg *config) error {
		if source == nil {
			return errors.New("token source must not be nil")
		}
		cfg.tokenSource = source
		return nil
	})
}

// WithUserAgent sets the User-Agent header value.
func WithUserAgent(userAgent string) Option {
	return optionFunc(func(cfg *config) error {
		if userAgent == "" {
			return errors.New("user-agent must not be empty")
		}
		cfg.userAgent = userAgent
		return nil
	})
}
