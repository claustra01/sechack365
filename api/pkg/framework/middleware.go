package framework

import (
	"net/http"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// TODO: save session to database or cache
type Session struct {
	Id        string
	UserId    string
	CreatedAt time.Time
	ExpiredAt time.Time
}

var Sessions = make(map[string]Session)

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

func CorsMiddleware(host string) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// well-known
			if strings.HasPrefix(r.URL.Path, "/.well-known") {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET")
				next(w, r)
				return
			}

			// api
			origin := r.Header.Get("Origin")
			if origin == host {
				w.Header().Set("Access-Control-Allow-Origin", host)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next(w, r)
		}
	}
}

func AuthMiddleware(logger model.ILogger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil || cookie.Value == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			sessionId := cookie.Value
			session, ok := Sessions[sessionId]
			if !ok || session.ExpiredAt.Before(time.Now()) {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}
}

func DevApiMiddleware(logger model.ILogger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Host, "localhost") {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			next(w, r)
		}
	}
}
