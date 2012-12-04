package mockhttp

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

type MockResponseWriter struct {
	Headers http.Header
	Body    bytes.Buffer
	Status  int
}

func (w *MockResponseWriter) Header() http.Header {
	return w.Headers
}

func (w *MockResponseWriter) Write(data []byte) (int, error) {
	return w.Body.Write(data)
}

func (w *MockResponseWriter) WriteHeader(status int) {
	w.Status = status
}

func (w *MockResponseWriter) Check(t *testing.T, want_status int, want_headers http.Header, want_body string) (ok bool) {
	ok = true
	if w.Status != want_status {
		t.Errorf("Bad HTTP status: got %d want %d", w.Status, want_status)
		ok = false
	}
	if !reflect.DeepEqual(w.Headers, want_headers) {
		t.Errorf("Bad HTTP response headers: %v", w.Headers)
		ok = false
	}
	resp_body := w.Body.String()
	if resp_body != want_body {
		t.Errorf("Bad HTTP response body: %q", resp_body)
		ok = false
	}
	return
}

func NewResponseWriter() *MockResponseWriter {
	var w MockResponseWriter
	w.Headers = make(http.Header)
	w.Status = 200
	return &w
}
