package framework

import (
	"fmt"
	"net/http"
)

var ErrRouteAlreadyExists = fmt.Errorf("route already exists")

type HandlerFunc func(*Config) http.HandlerFunc

type RouterInterface interface {
	Get(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error
	Post(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error
	Put(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error
	Delete(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error
	Patch(path string, handler HandlerFunc, middleware ...MiddlewareFunc) error
	Group(path string, middleware ...MiddlewareFunc) Router
	Use(middleware ...MiddlewareFunc)
}

type Router struct {
	Config     *Config
	mux        *http.ServeMux
	basePath   string
	middleware []MiddlewareFunc
}

func NewRouter(config *Config) *Router {
	r := &Router{
		Config:     config,
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
	r.mux.HandleFunc(pattern, chain(middleware...)(handler(r.Config)))
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
	router := *r
	router.basePath = r.basePath + path
	router.middleware = append(r.middleware, middleware...)
	return router
}

func (r *Router) Use(middleware ...MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware...)
}
