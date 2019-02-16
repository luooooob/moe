package minus

import "net/http"

// Routes is
type Routes interface {
	Set(method, pattern string, f http.HandlerFunc)
	Get(string, string) http.HandlerFunc
}

type defaultRoutes map[string]map[string]http.HandlerFunc

func newRoutes() defaultRoutes {
	return make(defaultRoutes)
}

// Set is
func (r defaultRoutes) Set(method, pattern string, f http.HandlerFunc) {
	r[method][pattern] = f
}

// Get is
func (r defaultRoutes) Get(method, pattern string) http.HandlerFunc {
	return r[method][pattern]
}
