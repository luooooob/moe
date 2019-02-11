package moe

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Request is
type Request struct {
	http.Request
}

// newContext returns a new Context instance for the given res and req
func newRequest(r *http.Request) *Request {
	return &Request{Request: *r}
}

// Path returns request path
func (req *Request) Path() string {
	return req.URL.Path
}

// Origin returns the Origin header
func (req *Request) Origin() string {
	return req.Header.Get("Origin")
}

// Type returns the Content-Type header
func (req *Request) Type(a *url.URL) string {
	return req.Header.Get("Content-Type")
}

// Bind binds the request body to the given pointer
func (req *Request) Bind(v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}

// Query is todo
func (req *Request) Query(key string) []string {
	return []string{""}
}

// IP is request remote address. Supports X-Forwarded-For if exists
func (req *Request) IP() string {
	ipString := req.Header.Get("X-Forwarded-For")
	if ipString != "" {
		IPs := strings.Split(ipString, ",")
		return IPs[0]
	}
	return req.RemoteAddr
}

// Secure is
func (req *Request) Secure() bool {
	return req.TLS != nil
}

// Fresh is  todo
func (req *Request) Fresh() bool {
	return false
}

// Idempotent is todo
func (req *Request) Idempotent() bool {
	return false
}
