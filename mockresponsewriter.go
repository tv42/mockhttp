package mockhttp

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// CheckResponse checks that the http response matches expectations,
// and calls t.Error if not.
func CheckResponse(t *testing.T, w *httptest.ResponseRecorder, wantStatus int, wantHeaders http.Header, wantBody string) (ok bool) {
	ok = true
	if w.Code != wantStatus {
		t.Errorf("Bad HTTP status: got %d want %d", w.Code, wantStatus)
		ok = false
	}
	if !reflect.DeepEqual(w.HeaderMap, wantHeaders) {
		t.Errorf("Bad HTTP response headers: %v", w.HeaderMap)
		ok = false
	}
	respBody := w.Body.String()
	if respBody != wantBody {
		t.Errorf("Bad HTTP response body: %q", respBody)
		ok = false
	}
	return
}
