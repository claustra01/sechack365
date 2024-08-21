package framework

import (
	"net/http"
)

type Server struct {
	Ctx    *Context
	Router *Router
	srv    *http.Server
}

func NewServer(ctx *Context) *Server {
	router := NewRouter(ctx)
	return &Server{
		srv: &http.Server{
			Addr:    ":" + ctx.Config.Port,
			Handler: router.mux,
		},
		Router: router,
	}
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}
