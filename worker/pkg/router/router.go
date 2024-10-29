package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	ServeMux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (router *Router) Get(path string, middleware func(next http.Handler) http.Handler, handler http.Handler) {
	pattern := fmt.Sprintf("GET %s", path)

	if middleware == nil {
		router.ServeMux.Handle(pattern, handler)
		return
	}
	router.ServeMux.Handle(pattern, middleware(handler))
}

func (router *Router) Post(path string, middleware func(next http.Handler) http.Handler, handler http.Handler) {
	pattern := fmt.Sprintf("POST %s", path)

	if middleware == nil {
		router.ServeMux.Handle(pattern, handler)
		return
	}
	router.ServeMux.Handle(pattern, middleware(handler))
}

func (router *Router) Put(path string, middleware func(next http.Handler) http.Handler, handler http.Handler) {
	pattern := fmt.Sprintf("PUT %s", path)

	if middleware == nil {
		router.ServeMux.Handle(pattern, handler)
		return
	}
	router.ServeMux.Handle(pattern, middleware(handler))
}

func (router *Router) Patch(path string, middleware func(next http.Handler) http.Handler, handler http.Handler) {
	pattern := fmt.Sprintf("PATCH %s", path)

	if middleware == nil {
		router.ServeMux.Handle(pattern, handler)
		return
	}
	router.ServeMux.Handle(pattern, middleware(handler))
}

func (router *Router) Delete(path string, middleware func(next http.Handler) http.Handler, handler http.Handler) {
	pattern := fmt.Sprintf("DELETE %s", path)

	if middleware == nil {
		router.ServeMux.Handle(pattern, handler)
		return
	}
	router.ServeMux.Handle(pattern, middleware(handler))
}
