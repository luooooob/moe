package moe

import (
	"encoding/json"
	"net/http"
)

// 1. response.Status()
// 2. response.Get()
// 3. response.Set()
// 4. response.Del()
// 6. response.Body()
// 7. response.Type()
// 8. response.ETag()
// 9. response.Vary(field)
// 10. response.LastModified
// 11. response.Flush()
// 12. response.Redirect(url)  // http.redirect

// Response is
type Response struct {
	Status int
	body   []byte
	Header http.Header

	LastModified string
}

func newResponse(w http.ResponseWriter) *Response {
	return &Response{}
}

// Send is
func (res *Response) Send(data interface{}) error {
	// c.ContentType("application/json")
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res.body = bytes
	return nil
}
