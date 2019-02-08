package moe

import (
	"container/list"
	"encoding/json"
	"log"
	"net/http"
)

const (
	JSONEncodeError = "JSON Encode Error"
)

// Context is
type Context struct {
	middlewareList *list.List
	writer         http.ResponseWriter

	Method        string
	Path          string
	UserAgent     string
	Authorization string
}

// newContext returns a new Context instance for the given res and req
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer: w,

		Method:        r.Method,
		Path:          r.URL.Path,
		UserAgent:     r.Header.Get("User-Agent"),
		Authorization: r.Header.Get("Authorization"),
	}
}

// setStatus is a wrapper for c.Writer.WriteHeader
// The provided code must be a valid HTTP 1xx-5xx status code
func (c *Context) Status(code int) *Context {
	c.writer.WriteHeader(code)
	return c
}

// setContentType is a wrapper for c.Writer.Header().Set("Content-Type", value)
func (c *Context) ContentType(value string) *Context {
	c.writer.Header().Set("Content-Type", value)
	return c
}

// JSON convert some value to JSON, and write them to response
func (c *Context) Send(code int, obj interface{}) *Context {
	c.ContentType("application/json")
	err := json.NewEncoder(c.writer).Encode(obj)
	if err != nil {
		log.Println(err)
		return c.Error(500, JSONEncodeError)
	}
	return c.Status(code)
}

// // Assert is similar to .Error(), it Errors when:
// // statement(error type) != nil or
// // statement(bool) == false`
// func (c *Context) Assert(assertion interface{}, code int, msg string) *Context {
// 	switch a := assertion.(type) {
// 	case error:
// 		if a != nil {
// 			return c.Error(code, msg)
// 		}
// 	case bool:
// 		if !a {
// 			return c.Error(code, msg)
// 		}
// 	}
// 	return c
// }

// Error to formart Error message
type Error struct {
	Message string `json:"message"`
}

// Throw is like http.Error(), convert error msg to json form
func (c *Context) Error(code int, msg string) *Context {
	c.Status(code)
	if err := json.NewEncoder(c.writer).Encode(&Error{msg}); err != nil {
		http.Error(c.Writer, msg, code)
	}
	return c
}

// Redirect is a wrapper for http.Redirect
func (c *Context) Redirect(code int, url string) *Context {
	http.Redirect(c.Writer, c.Request, url, code)
	return c
}

// Next is
func (c *Context) Next() *Context {
	ele := c.middlewareList.Remove(c.middlewareList.Front())
	if f, ok := ele.(Controller); ok {
		f(c)
	}
	return c
}
