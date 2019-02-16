package minus

import (
	"net/http"
)

// Adapter is
type Adapter func(http.Handler) http.Handler

// App is
type App struct {
	handler http.Handler
}

// MiddlewareAdapter is
type MiddlewareAdapter struct {
	MiddlewareFunc http.HandlerFunc
}

// ServHTTP is
func (m *MiddlewareAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.MiddlewareFunc(w, r)
}

// NewApp returns a new Moe object
func NewApp() *App {
	return &App{}
}

// ServeHTTP is
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}

// Use is
func (a *App) Use(func(http.ResponseWriter, *http.Request, http.HandlerFunc)) {
	Adapter := &MiddlewareAdapter{}
	Adapter.MiddlewareFunc = func(http.ResponseWriter, *http.Request) {

	}
	return
}

func example() {
	a := NewApp()
	a.Use(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		next(w, r)
	})
}
