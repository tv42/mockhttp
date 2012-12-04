package mockhttp_test

import (
	"github.com/tv42/mockhttp.go"
	"net/http"
	"strings"
	"testing"
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
	req := mockhttp.NewRequest("GET", "http://foo.example.com/bar", nil)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	respw := mockhttp.NewResponseWriter()
	handler.ServeHTTP(respw, req)
	want_hdr := make(http.Header)
	want_hdr.Add("Content-Type", "text/plain; charset=utf-8")
	respw.Check(t, http.StatusOK, want_hdr, "Hello, world.\n")
}

func TestPUT(t *testing.T) {
	handler := http.HandlerFunc(hello)
	body := strings.NewReader(`foo`)
	req := mockhttp.NewRequest("PUT", "http://foo.example.com/bar", body)
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	respw := mockhttp.NewResponseWriter()
	handler.ServeHTTP(respw, req)
	want_hdr := make(http.Header)
	want_hdr.Add("Content-Type", "text/plain; charset=utf-8")
	want_hdr.Add("Allow", "GET")
	respw.Check(t, http.StatusMethodNotAllowed, want_hdr, "Only GET supported.\n")
}
