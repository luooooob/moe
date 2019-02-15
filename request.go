package moe

import "net/http"

// Request is
type Request struct {
	http.Request
}

func newRequest(r *http.Request) *Request {
	return &Request{*r}
}
