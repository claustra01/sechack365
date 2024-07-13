package framework

import (
	"log/slog"
	"net/http"
)

type MiddlewareFunc func(HandlerFunc) HandlerFunc

func chain(middleware ...MiddlewareFunc) MiddlewareFunc {
	return func(handler HandlerFunc) HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			handler = middleware[i](handler)
		}
		return handler
	}
}

func LoggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request:", "Method", r.Method, "Path", r.URL.Path, "RemoteAddr", r.RemoteAddr, "Proto", r.Proto, "UserAgent", r.UserAgent())
		next(w, r)
	}
}

func RecoverMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Panic Recovered:", "Error", err.(string))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
