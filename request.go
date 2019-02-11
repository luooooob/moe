package moe

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Request is
type Request struct {
	http.Request

	Path       string
	Origin     string
	IP         string
	Type       string
	Fresh      bool
	Secure     bool
	Idempotent bool
}

// newContext returns a new Context instance for the given res and req
func newRequest(r *http.Request) *Request {
	req := &Request{Request: *r}
	req.Path = req.URL.Path
	req.Origin = req.Get("Origin")
	req.Type = req.Get("Content-Type")
	req.IP = req.getClientIP()
	req.Fresh = req.isFresh()
	req.Idempotent = req.isIdempotent()
	return req
}

// Get is
func (req *Request) Get(key string) string {
	return req.Header.Get(key)
}

// Parse is
func (req *Request) Parse(v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
}

// Query is todo
func (req *Request) Query(key string) string {
	return ""
}

// getClientIP is
func (req *Request) getClientIP() string {
	ipString := req.Get("X-Forwarded-For")
	if ipString != "" {
		IPs := strings.Split(ipString, ",")
		return IPs[0]
	}
	return req.RemoteAddr
}

// isSecure is
func (req *Request) isSecure() bool {
	return req.TLS != nil
}

// isFresh is  todo
func (req *Request) isFresh() bool {
	return false
}

// isIdempotent is todo
func (req *Request) isIdempotent() bool {
	return false
}
