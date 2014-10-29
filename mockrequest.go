package mockhttp

import (
	"io"
	"net/http"
	"testing"
)

// NewRequest is like http.NewRequest, but calls t.Fatal on error.
func NewRequest(t testing.TB, method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("Bug in test: cannot construct http.Request from method=%q, url=%q, body=%#v: %s", method, url, body, err)
	}
	return req
}
