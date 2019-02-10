package moe

import (
	"encoding/json"
	"io"
	"net/http"
)

// Request is
type Request struct {
	Method string
	Path   string
	Header http.Header
	body   io.Reader
}

// newContext returns a new Context instance for the given res and req
func newRequest(r *http.Request) *Request {
	return &Request{
		Method: r.Method,
		Path:   r.URL.Path,
		Header: r.Header,
		body:   r.Body,
	}
}

// Parse is
func (req *Request) Parse(v interface{}) error {
	return json.NewDecoder(req.body).Decode(v)
}

// Get is
func (req *Request) Get(key string) string {
	return req.Header.Get(key)
}
