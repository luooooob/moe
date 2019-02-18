package moe

import (
	"log"
	"net/http"
)

// Controller is
type Controller func(w http.ResponseWriter, r *http.Request, next Next)

// MiddlewareList is
type MiddlewareList interface {
	Add(middleware)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type defaultMiddlewareList struct {
	middlewares []middleware
	handlerFunc *Controller
}

var (
	_ MiddlewareList = defaultMiddlewareList{}
	_ http.Handler   = defaultMiddlewareList{}
)

type middleware Controller

// Next is
type Next func()

func m1(w http.ResponseWriter, r *http.Request, next Next) {
	// do something
	log.Println("before1")
	next()
	log.Println("after4")
}

func m2(w http.ResponseWriter, r *http.Request, next Next) {

	log.Println("before2")
	next()
	log.Println("after3")
}

// comppose is
func compose(m1, m2 middleware) middleware {
	return func(w http.ResponseWriter, r *http.Request, next Next) {
		m1(w, r, func() {
			m2(w, r, next)
		})
	}
}

func newMiddlewareList() *defaultMiddlewareList {
	return &defaultMiddlewareList{
		middlewares: make([]middleware, 0),
	}
}

func (l defaultMiddlewareList) Add(middleware middleware) {
	// handlerFunc := l.handlerFunc

	// *l.handlerFunc = func(w http.ResponseWriter, r *http.Request, next *Controller) error {
	// 	next = handlerFunc
	// 	return middleware(w, r, next)
	// }
}

func (l defaultMiddlewareList) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
