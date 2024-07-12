package framework

import (
	"fmt"
	"net/http"
)

var ErrRouteAlreadyExists = fmt.Errorf("route already exists")

type HandlerFunc http.HandlerFunc
type Middleware func(HandlerFunc) HandlerFunc

type RouterInterface interface {
	Get(path string, handler HandlerFunc, middleware ...Middleware) error
	Post(path string, handler HandlerFunc, middleware ...Middleware) error
	Put(path string, handler HandlerFunc, middleware ...Middleware) error
	Delete(path string, handler HandlerFunc, middleware ...Middleware) error
	Patch(path string, handler HandlerFunc, middleware ...Middleware) error
	HandleRoutes()
}

type Router struct {
	RouterInterface
	mux    *http.ServeMux
	routes map[string]map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		mux:    http.NewServeMux(),
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (r *Router) addRoute(path string, method string, handler HandlerFunc, middleware ...Middleware) error {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]HandlerFunc)
	}

	if r.routes[path][method] != nil {
		return ErrRouteAlreadyExists
	}

	finalHandler := handler
	for _, m := range middleware {
		finalHandler = m(finalHandler)
	}
	r.routes[path][method] = finalHandler
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

func (r *Router) HandleRoutes() {
	for path, handlers := range r.routes {
		r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
			handler := handlers[req.Method]
			if handler == nil {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			handler(w, req)
		})
	}
}
