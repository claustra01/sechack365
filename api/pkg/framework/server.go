package framework

import (
	"net/http"
)

type ServerInterface interface {
	ListenAndServe() error
}

type Server struct {
	ServerInterface
	router Router
	port   string
}

func NewServer(router Router, port string) *Server {
	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(":"+s.port, s.router.mux)
}
