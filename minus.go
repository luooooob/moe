package minus

import (
	"net/http"
	"path"
)

// App is
type App struct {
	prefix      string
	middlewares []middleware
	router      Router
}

type routerMap map[string]map[string]http.HandlerFunc

type middleware func(value interface{}) http.HandlerFunc

// newApp make and return a App
func newApp(prefix string) *App {
	return &App{
		prefix:      prefix,
		middlewares: make([]middleware, 0),
		routeMap:    newRouter()
	}
}

// Register registers the http.HandlerFunc for the given method and pattern
func (router *App) Register(method, pattern string, c http.HandlerFunc) *App {
	pattern = path.Join(r.prefix, pattern)
	r.routeMap[method][pattern] = c
	return r
}

// Use is
func (router *App) Use(m middleware) *App {
	r.middlewares = append(r.middlewares, m)
	return r
}

func (router *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
