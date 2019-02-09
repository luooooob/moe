package moe

import (
	"net/http"
)

const null = ""

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
func (m *Moe) register(method, pattern string, f Poi) *Moe {
	m.routes[method][pattern] = f
	return m
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

// // GET is a shortcut for Handle("GET", pattern, f)
// func (m *Moe) GET(pattern string, f Poi) *Moe {
// 	return m.register("GET", pattern, f)
// }

// // POST is a shortcut for Handle("POST", pattern, f)
// func (m *Moe) POST(pattern string, f Poi) *Moe {
// 	return m.register("POST", pattern, f)
// }

// // PUT is a shortcut for Handle("PUT", pattern, f)
// func (m *Moe) PUT(pattern string, f Poi) *Moe {
// 	return m.register("PUT", pattern, f)
// }

// // DELETE is a shortcut for Handle("DELETE", pattern, f)
// func (m *Moe) DELETE(pattern string, f Poi) *Moe {
// 	return m.register("DELETE", pattern, f)
// }
