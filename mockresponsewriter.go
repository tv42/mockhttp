package mockhttp

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func CheckResponse(t *testing.T, w *httptest.ResponseRecorder, want_status int, want_headers http.Header, want_body string) (ok bool) {
	ok = true
	if w.Code != want_status {
		t.Errorf("Bad HTTP status: got %d want %d", w.Code, want_status)
		ok = false
	}
	if !reflect.DeepEqual(w.HeaderMap, want_headers) {
		t.Errorf("Bad HTTP response headers: %v", w.HeaderMap)
		ok = false
	}
	resp_body := w.Body.String()
	if resp_body != want_body {
		t.Errorf("Bad HTTP response body: %q", resp_body)
		ok = false
	}
	return
}
