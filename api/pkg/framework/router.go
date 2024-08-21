package framework

import (
	"fmt"
	"net/http"
)

var ErrRouteAlreadyExists = fmt.Errorf("route already exists")

type HandlerFunc func(*Context) http.HandlerFunc

type Router struct {
	Ctx        *Context
	mux        *http.ServeMux
	basePath   string
	middleware []MiddlewareFunc
}

func NewRouter(ctx *Context) *Router {
	r := &Router{
		Ctx:        ctx,
		mux:        http.NewServeMux(),
		basePath:   "",
		middleware: []MiddlewareFunc{},
	}
	r.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	return r
}

func (r *Router) addRoute(path string, method string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	p := r.basePath + path
	if p[len(p)-1] == '/' {
		p += "{$}"
	}

	pattern := fmt.Sprintf("%s %s", method, p)
	middleware = append(r.middleware, middleware...)
	r.mux.HandleFunc(pattern, chain(middleware...)(handler(r.Ctx)))
	return nil
}

func (r *Router) Get(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	return r.addRoute(path, http.MethodGet, handler, middleware...)
}

func (r *Router) Post(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	return r.addRoute(path, http.MethodPost, handler, middleware...)
}

func (r *Router) Put(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	return r.addRoute(path, http.MethodPut, handler, middleware...)
}

func (r *Router) Delete(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	return r.addRoute(path, http.MethodDelete, handler, middleware...)
}

func (r *Router) Patch(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error {
	return r.addRoute(path, http.MethodPatch, handler, middleware...)
}

func (r *Router) Group(path string, middleware ...MiddlewareFunc) Router {
	return Router{
		Ctx:        r.Ctx,
		mux:        r.mux,
		basePath:   r.basePath + path,
		middleware: append(r.middleware, middleware...),
	}
}

func (r *Router) Use(middleware ...MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware...)
}
