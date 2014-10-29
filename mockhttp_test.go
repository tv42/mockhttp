package mockhttp_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tv42/mockhttp"
)

func hello(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.Header().Add("Allow", "GET")
		http.Error(w, "Only GET supported.", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world.\n"))
}

func TestGET(t *testing.T) {
	handler := http.HandlerFunc(hello)
	req := mockhttp.NewRequest(t, "GET", "http://foo.example.com/bar", nil)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	respw := httptest.NewRecorder()
	handler.ServeHTTP(respw, req)
	wantHdr := make(http.Header)
	wantHdr.Add("Content-Type", "text/plain; charset=utf-8")
	mockhttp.CheckResponse(t, respw, http.StatusOK, wantHdr, "Hello, world.\n")
}

func TestPUT(t *testing.T) {
	handler := http.HandlerFunc(hello)
	body := strings.NewReader(`foo`)
	req := mockhttp.NewRequest(t, "PUT", "http://foo.example.com/bar", body)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	respw := httptest.NewRecorder()
	handler.ServeHTTP(respw, req)
	wantHdr := make(http.Header)
	wantHdr.Add("Content-Type", "text/plain; charset=utf-8")
	wantHdr.Add("Allow", "GET")
	mockhttp.CheckResponse(t, respw, http.StatusMethodNotAllowed, wantHdr, "Only GET supported.\n")
}
