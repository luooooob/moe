package moe

import (
	"container/list"
	"encoding/json"
	"net/http"
)

// Context is
type Context struct {
	middlewareList *list.List
	Writer         http.ResponseWriter
	Method         string
	Path           string
}

// newContext returns a new Context instance for the given res and req
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}

// // Set is used to store a new key/value pair exclusively for this context.
// // use golang context
// func (c *Context) Set(key, value interface{}) *Context {
// 	c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
// 	return c
// }

// // Get returns the value for the given key
// // use golang context
// func (c *Context) Get(key interface{}) interface{} {
// 	return c.Request.Context().Value(key)
// }

// Method returns the http request method as string
// ex: "GET"
// func (c *Context) Method() string {
// 	return c.Request.Method
// }

// Path is
func (c *Context) Path() string {
	return c.Request.URL.Path
}

// Header returns the header for the given key
func (c *Context) Header(key string) string {
	return c.Request.Header.Get(key)
}

// UserAgent is a shortcut for c.Header("User-Agent")
func (c *Context) UserAgent() string {
	return c.Header("User-Agent")
}

// Authorization is a shortcut for c.Header("Authorization")
func (c *Context) Authorization() string {
	return c.Header("Authorization")
}

// ParseForm is
func (c *Context) ParseForm() *Context {
	return c.Assert(c.Request.ParseForm(), http.StatusBadRequest, "")
}

// Form is a wrapper for req.Form.Get()
func (c *Context) Form(key string) string {
	return c.Request.Form.Get(key)
}

// Status is a wrapper for c.Writer.WriteHeader
// The provided code must be a valid HTTP 1xx-5xx status code
func (c *Context) Status(code int) *Context {
	c.Writer.WriteHeader(code)
	return c
}

// ContentType is a shortcut for c.Writer.Header().Set("Content-Type", value)
func (c *Context) ContentType(value string) *Context {
	c.Writer.Header().Set("Content-Type", value)
	return c
}

// Write is an alias for c.res.Write(body)
func (c *Context) Write(body []byte) *Context {
	c.Writer.Write(body)
	return c
}

// JSON convert some value to JSON, and write them to response
func (c *Context) JSON(code int, body interface{}) *Context {
	err := json.NewEncoder(c.Writer).Encode(body)
	return c.
		Status(code).
		ContentType("application/json").
		Assert(err, 500, "JSON Encode Error")
}

// Assert is similar to .Error(), it Errors when:
// statement(error type) != nil or
// statement(bool) == false`
func (c *Context) Assert(assertion interface{}, code int, msg string) *Context {
	switch a := assertion.(type) {
	case error:
		if a != nil {
			return c.Error(code, msg)
		}
	case bool:
		if !a {
			return c.Error(code, msg)
		}
	}
	return c
}

// Error to formart Error message
type Error struct {
	Message string `json:"message"`
}

// Throw is like http.Error(), convert error msg to json form
func (c *Context) Error(code int, msg string) *Context {
	c.Status(code)
	if err := json.NewEncoder(c.Writer).Encode(&Error{msg}); err != nil {
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
