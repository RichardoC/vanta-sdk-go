package v1

import (
	"net/http"
	"strings"
	"testing"
)

func TestAPIErrorString(t *testing.T) {
	err := &APIError{StatusCode: http.StatusBadRequest, Status: "400 Bad Request", Body: []byte(`{"error":"invalid"}`)}
	msg := err.Error()
	if !strings.Contains(msg, "400") || !strings.Contains(msg, "invalid") {
		t.Fatalf("unexpected error string: %s", msg)
	}
}

func TestDecodeEmptyRaw(t *testing.T) {
	type sample struct{ Name string }
	out, err := Decode[sample](nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Name != "" {
		t.Fatalf("unexpected decoded value: %+v", out)
	}
}

func TestDecodeInvalidRaw(t *testing.T) {
	type sample struct{ Name string }
	_, err := Decode[sample]([]byte("not-json"))
	if err == nil {
		t.Fatal("expected decode error")
	}
	if !strings.Contains(err.Error(), "decode JSON payload") {
		t.Fatalf("unexpected error: %v", err)
	}
}
