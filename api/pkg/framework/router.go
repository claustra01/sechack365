package framework

import (
	"fmt"
	"net/http"
)

var ErrRouteAlreadyExists = fmt.Errorf("route already exists")

type HandlerFunc http.HandlerFunc

type RouterInterface interface {
	Get(path string, handler HandlerFunc, middleware ...Middleware) error
	Post(path string, handler HandlerFunc, middleware ...Middleware) error
	Put(path string, handler HandlerFunc, middleware ...Middleware) error
	Delete(path string, handler HandlerFunc, middleware ...Middleware) error
	Patch(path string, handler HandlerFunc, middleware ...Middleware) error
	Group(path string, middleware ...Middleware) Router
	Use(middleware ...Middleware)
}

type Router struct {
	mux        *http.ServeMux
	basePath   string
	routes     map[string]map[string]HandlerFunc
	middleware []Middleware
}

func NewRouter() *Router {
	r := &Router{
		mux:        http.NewServeMux(),
		basePath:   "",
		routes:     make(map[string]map[string]HandlerFunc),
		middleware: []Middleware{},
	}
	r.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	return r
}

func (r *Router) addRoute(path string, method string, handler HandlerFunc, middleware ...Middleware) error {
	p := r.basePath + path
	if p[len(p)-1] == '/' {
		p += "{$}"
	}

	pattern := fmt.Sprintf("%s %s", method, p)
	middleware = append(r.middleware, middleware...)
	r.mux.HandleFunc(pattern, chain(middleware...)(handler))
	return nil
}

func (r *Router) Get(path string, handler HandlerFunc, middleware ...Middleware) error {
	return r.addRoute(path, http.MethodGet, handler, middleware...)
}

func (r *Router) Post(path string, handler HandlerFunc, middleware ...Middleware) error {
	return r.addRoute(path, http.MethodPost, handler, middleware...)
}

func (r *Router) Put(path string, handler HandlerFunc, middleware ...Middleware) error {
	return r.addRoute(path, http.MethodPut, handler, middleware...)
}

func (r *Router) Delete(path string, handler HandlerFunc, middleware ...Middleware) error {
	return r.addRoute(path, http.MethodDelete, handler, middleware...)
}

func (r *Router) Patch(path string, handler HandlerFunc, middleware ...Middleware) error {
	return r.addRoute(path, http.MethodPatch, handler, middleware...)
}

func (r *Router) Group(path string, middleware ...Middleware) Router {
	return Router{
		mux:        r.mux,
		basePath:   r.basePath + path,
		routes:     r.routes,
		middleware: append(r.middleware, middleware...),
	}
}

func (r *Router) Use(middleware ...Middleware) {
	r.middleware = append(r.middleware, middleware...)
}
