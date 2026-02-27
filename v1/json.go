package v1

import (
	"encoding/json"
	"fmt"
)

// Decode unmarshals a raw generated response body into a typed Go struct.
func Decode[T any](raw json.RawMessage) (T, error) {
	var out T
	if len(raw) == 0 {
		return out, nil
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return out, fmt.Errorf("decode JSON payload: %w", err)
	}
	return out, nil
}
