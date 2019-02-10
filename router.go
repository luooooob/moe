package moe

import (
	"net/http"
)

// Poi is the function can be registered
// note: Poi and HandleFunc are different
type Poi func(*Context)

type routes map[string]map[string]Poi

// newRoutes make and return a routes
func newRoutes() routes {
	return make(routes)
}

// Match registers the f for the given method
// and pattern with Composed Middlewares
func (m *Moe) register(method, pattern string, p Poi) {
	m.routes[method][pattern] = p
}

func (m *Moe) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	if f, ok := m.routes[c.Method][c.Path]; ok {
		middlewareList := *m.middlewareList
		middlewareList.PushBack(f)
		c.middlewareList = &middlewareList
		c.Next()
	} else {
		// c.Error(http.StatusNotFound, null)
	}
}

// func matchPath(method, path string) f {
// 	return
// }

// GET is a shortcut for Handle("GET", pattern, f)
func (m *Moe) GET(pattern string, p Poi) {
	m.register("GET", pattern, p)
}

// POST is a shortcut for Handle("POST", pattern, p)
func (m *Moe) POST(pattern string, p Poi) {
	m.register("POST", pattern, p)
}

// PUT is a shortcut for Handle("PUT", pattern, p)
func (m *Moe) PUT(pattern string, p Poi) {
	m.register("PUT", pattern, p)
}

// DELETE is a shortcut for Handle("DELETE", pattern, p)
func (m *Moe) DELETE(pattern string, p Poi) {
	m.register("DELETE", pattern, p)
}
