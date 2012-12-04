package mockhttp

import (
	"fmt"
	"io"
	"net/http"
)

// This is like http.NewRequest, but panics on errors.
func NewRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(fmt.Sprintf("Bug in test: cannot construct http.Request from method=%q, url=%q, body=%#v: %s", method, url, body, err))
	}
	return req
}
