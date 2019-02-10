package moe

import "net/http"

// Response is
type Response struct {
	status int
	body   []byte
	Header http.Header
	// LastModified string
}

func newResponse(w http.ResponseWriter) *Response {
	return &Response{}
}

// Set is
func (res *Response) Set(key, value string) {
	res.Header.Set(key, value)
}
