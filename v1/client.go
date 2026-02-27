package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultAPIBaseURL   = "https://api.vanta.com"
	defaultOAuthPath    = "/oauth/token"
	defaultUserAgent    = "vanta-sdk-go/0.1"
	contentTypeJSON     = "application/json"
	headerAuthorization = "Authorization"
)

// Client is the entrypoint for calling Vanta APIs.
type Client struct {
	httpClient  *http.Client
	baseURL     *url.URL
	tokenSource TokenSource
	userAgent   string

	// Generated service handles are populated by newGeneratedServices.
	Services *Services
}

// NewClient creates a configured API client.
func NewClient(opts ...Option) (*Client, error) {
	cfg := defaultConfig()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt.apply(cfg); err != nil {
			return nil, err
		}
	}

	baseURL, err := url.Parse(strings.TrimRight(cfg.baseURL, "/"))
	if err != nil {
		return nil, fmt.Errorf("parse base URL: %w", err)
	}

	c := &Client{
		httpClient:  cfg.httpClient,
		baseURL:     baseURL,
		tokenSource: cfg.tokenSource,
		userAgent:   cfg.userAgent,
	}
	c.Services = newGeneratedServices(c)
	return c, nil
}

// BaseURL returns the current API base URL.
func (c *Client) BaseURL() string {
	if c == nil || c.baseURL == nil {
		return ""
	}
	return c.baseURL.String()
}

func (c *Client) newRequest(ctx context.Context, method, path string, query url.Values, body any) (*http.Request, error) {
	u := *c.baseURL
	u.Path = joinURLPath(c.baseURL.Path, path)
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		buf := &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, fmt.Errorf("encode request JSON: %w", err)
		}
		bodyReader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Accept", contentTypeJSON)
	req.Header.Set("User-Agent", c.userAgent)
	if body != nil {
		req.Header.Set("Content-Type", contentTypeJSON)
	}

	if c.tokenSource != nil {
		tok, err := c.tokenSource.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("resolve bearer token: %w", err)
		}
		req.Header.Set(headerAuthorization, "Bearer "+tok.AccessToken)
	}

	return req, nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, decodeAPIError(resp)
	}
	return resp, nil
}

func (c *Client) doJSON(req *http.Request, out any) error {
	resp, err := c.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if out == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		if err == io.EOF {
			return nil
		}
		return fmt.Errorf("decode response JSON: %w", err)
	}
	return nil
}

func (c *Client) newMultipartRequest(ctx context.Context, method, path string, query url.Values, form map[string]string) (*http.Request, error) {
	u := *c.baseURL
	u.Path = joinURLPath(c.baseURL.Path, path)
	if len(query) > 0 {
		u.RawQuery = query.Encode()
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, value := range form {
		if err := writer.WriteField(key, value); err != nil {
			return nil, fmt.Errorf("write multipart field %q: %w", key, err)
		}
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("finalize multipart body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("build multipart request: %w", err)
	}
	req.Header.Set("Accept", contentTypeJSON)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", c.userAgent)

	if c.tokenSource != nil {
		tok, err := c.tokenSource.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("resolve bearer token: %w", err)
		}
		req.Header.Set(headerAuthorization, "Bearer "+tok.AccessToken)
	}

	return req, nil
}

func joinURLPath(basePath, relativePath string) string {
	basePath = strings.TrimSuffix(basePath, "/")
	relativePath = strings.TrimPrefix(relativePath, "/")
	if basePath == "" {
		return "/" + relativePath
	}
	if relativePath == "" {
		return basePath
	}
	return basePath + "/" + relativePath
}
