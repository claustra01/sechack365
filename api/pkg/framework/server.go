package framework

import (
	"net/http"
)

type ServerInterface interface {
	ListenAndServe() error
}

type Server http.Server

func NewServer(router Router, port string) *Server {
	return &Server{
		Addr:    ":" + port,
		Handler: router.mux,
	}
}

func (s *Server) ListenAndServe() error {
	return (*http.Server)(s).ListenAndServe()
}
