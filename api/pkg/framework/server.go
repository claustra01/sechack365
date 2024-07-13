package framework

import (
	"log/slog"
	"net/http"
)

type ServerInterface interface {
	ListenAndServe() error
}

type Server struct {
	Ctx    *Context
	Router *Router
	srv    *http.Server
}

func NewServer(ctx *Context) *Server {
	router := NewRouter(ctx)

	slog.Info("LogLevel set to:", "level", ctx.Config.LogLevel)
	slog.SetLogLoggerLevel(ctx.Config.LogLevel)

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
