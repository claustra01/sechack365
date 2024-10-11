package framework

import (
	"log"
	"net/http"
	"strings"

	"github.com/claustra01/sechack365/pkg/model"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func chain(middleware ...MiddlewareFunc) MiddlewareFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			handler = middleware[i](handler)
		}
		return handler
	}
}

func LoggingMiddleware(logger model.ILogger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Middleware Logging", "Method", r.Method, "Path", r.URL.Path, "RemoteAddr", r.RemoteAddr, "Proto", r.Proto, "UserAgent", r.UserAgent())
			next(w, r)
		}
	}
}

func RecoverMiddleware(logger model.ILogger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("Panic Recovered", "Error", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next(w, r)
		}
	}
}

func DevApiMiddleware(logger model.ILogger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Host)
			if !strings.HasPrefix(r.Host, "localhost") {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			next(w, r)
		}
	}
}
