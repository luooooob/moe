package moe

import (
	"container/list"
)

// Use appends a Poi func to the HandlerFunc chain.
// func (m *Moe) Use(mws ...Middleware) *Moe {
// 	for _, mw := range mws {
// 		m.mws = append(m.mws, mw)
// 	}
// 	return m
// }

func newMiddlewareList() *list.List {
	return list.New()
}

// Use is
func (m *Moe) Use(f Poi) *Moe {
	m.middlewareList.PushBack(f)
	return m
}
