package moe

import (
	"container/list"
)

// Moe is justice!
type Moe struct {
	middlewareList *list.List

	routes map[string]map[string]Controller
}

// NewApp returns a new Moe instance
func NewApp() *Moe {
	return &Moe{
		routes:         newRoutes(),
		middlewareList: newMiddlewareList(),
	}
}

// // Listen is a wrapper for http.ListenAndServe
// func (m *Moe) Listen(addr string) *Moe {
// 	http.ListenAndServe(addr, m.mux)
// 	return m
// }
