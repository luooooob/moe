package moe

import (
	"container/list"
	"encoding/json"
	"net/http"
)

const (
	jsonEncodeError = "JSON Encode Error"
	jsonContentType = "application/json"
)

// Context is
type Context struct {
	w http.ResponseWriter
	r *http.Request

	middlewareList *list.List

	Request  Request
	Response Response

	Method        string
	Path          string
	Header        http.Header
	UserAgent     string
	Authorization string
}

// newContext returns a new Context instance for the given res and req
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w: w,
		r: r,

		Method:        r.Method,
		Path:          r.URL.Path,
		Header:        r.Header,
		UserAgent:     r.Header.Get("User-Agent"),
		Authorization: r.Header.Get("Authorization"),
	}
}

// setHeader is
func (c *Context) setHeader(key, value string) {
	c.w.Header().Set(key, value)
}

func (c *Context) setAuthorization(value string) {
	c.setHeader("Authorization", value)
}

// setContentType is a wrapper for c.Writer.Header().Set("Content-Type", value)
func (c *Context) setContentType(value string) {
	c.setHeader("Content-Type", value)
}

// setStatus is a wrapper for c.Writer.WriteHeader
// The provided code must be a valid HTTP 1xx-5xx status code
func (c *Context) setStatus(code int) {
	c.w.WriteHeader(code)
}

// Send convert some value to JSON, and write them to response
func (c *Context) Send(value interface{}) error {
	// c.ContentType("application/json")
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.write(jsonBytes)
}

func (c *Context) write(b []byte) error {
	_, err := c.w.Write(b)
	if err != nil {
		panic(err)
	}
	return nil
}

// Redirect is a wrapper for http.Redirect
func (c *Context) Redirect(code int, url string) *Context {
	http.Redirect(c.w, c.r, url, code)
	return c
}

// Next is
func (c *Context) Next() *Context {
	ele := c.middlewareList.Remove(c.middlewareList.Front())
	if f, ok := ele.(Poi); ok {
		f(c)
	}
	return c
}
