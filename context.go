package moe

import "net/http"

// Context is
type Context struct {
	// middlewareList *list.List

	Request  *Request
	Response *Response
}

// newContext returns a new Context instance for the given res and req
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Request:  newRequest(r),
		Response: newResponse(w),
	}
}

// writeResponse is
func (c *Context) writeResponse(w http.ResponseWriter) error {
	_, err := w.Write(c.Response.body)
	if err != nil {
		panic(err)
	}
	return nil
}

// writeBody is
func (c *Context) writeBody(w http.ResponseWriter) error {
	_, err := w.Write(c.Response.body)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *Context) writeHeader(w http.ResponseWriter) {

}
