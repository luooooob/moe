package moe

import "net/http"

// Routes is
type Routes interface {
	Set(method, pattern string, f http.HandlerFunc)
	Get(string, string) http.HandlerFunc
}

type defaultRoutes struct {
	routemap routemap
}

type routemap map[string]map[string]http.HandlerFunc

func newRoutes() *defaultRoutes {
	return &defaultRoutes{
		routemap: make(routemap),
	}
}

// Set is
func (r defaultRoutes) Set(method, pattern string, f http.HandlerFunc) {
	r.routemap[method][pattern] = f
}

// Get is
func (r defaultRoutes) Get(method, pattern string) http.HandlerFunc {
	return r.routemap[method][pattern]
}
