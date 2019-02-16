package minus

import (
	"net/http"
)

// App is
type App struct {
	prefix         string
	routes         Routes
	middlewareList MiddlewareList
}

// newApp make and return a App
func newApp(prefix string) *App {
	return &App{
		prefix:         prefix,
		routes:         newRoutes(),
		middlewareList: newMiddlewareList(),
	}
}

// Register registers the http.HandlerFunc for the given method and pattern
func (a *App) Register(method, pattern string, f http.HandlerFunc) *App {
	a.routes.Set(method, pattern, f)
	return a
}

// Use is
func (a *App) Use(m middleware) *App {
	a.middlewareList.Add(m)
	return a
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
