package moe

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Request is derived from http.Request and adding some methods
type Request http.Request

// newContext returns a new Request instance with the given *http.Request
func newRequest(r *http.Request) *Request {
	return (*Request)(r)
}

// Path returns request path
func (req *Request) Path() string {
	return req.URL.Path
}

// Type returns the Content-Type header
func (req *Request) Type(a *url.URL) string {
	return req.Header.Get("Content-Type")
}

// Origin returns the Origin header
func (req *Request) Origin() string {
	return req.Header.Get("Origin")
}

// Bind binds the request body with the given pointer
func (req *Request) Bind(v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}

// Query returns the first value with the given key in querystring
func (req *Request) Query(key string) string {
	return req.URL.Query().Get(key)
}

// QueryAll is todo
func (req *Request) QueryAll(key string) []string {
	return req.URL.Query()[key]
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

// Fresh is  todo
func (req *Request) Fresh() bool {
	return false
}

// Idempotent is todo
func (req *Request) Idempotent() bool {
	return false
}
