package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIError wraps non-2xx API responses.
type APIError struct {
	StatusCode int
	Status     string
	Body       []byte
	ParsedBody map[string]any
}

func (e *APIError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if len(e.Body) > 0 {
		return fmt.Sprintf("vanta API error: %d %s: %s", e.StatusCode, e.Status, truncate(e.Body, 512))
	}
	return fmt.Sprintf("vanta API error: %d %s", e.StatusCode, e.Status)
}

func decodeAPIError(resp *http.Response) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 64*1024))
	err := &APIError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       body,
	}
	if len(body) > 0 {
		var parsed map[string]any
		if json.Unmarshal(body, &parsed) == nil {
			err.ParsedBody = parsed
		}
	}
	return err
}

func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "..."
}
