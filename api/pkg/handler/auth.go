package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

func Login(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequestBody openapi.Auth
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &authRequestBody)
		if err != nil {
			returnBadRequest(w, c.Logger, err)
			return
		}

		user, err := c.Controllers.User.FindByUsername(authRequestBody.Username, c.Config.Host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if user == nil {
			returnUnauthorized(w, c.Logger, nil)
			return
		}
		if err := util.CompareHashAndPassword(user.HashedPassword, authRequestBody.Password); err != nil {
			returnUnauthorized(w, c.Logger, err)
			return
		}

		for _, session := range framework.Sessions {
			if session.UserId == user.Id {
				delete(framework.Sessions, session.Id)
				break
			}
		}

		sessionId := util.NewUuid().String()
		framework.Sessions[sessionId] = framework.Session{
			Id:        sessionId,
			UserId:    user.Id,
			CreatedAt: time.Now(),
			ExpiredAt: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionId,
			HttpOnly: true,
			Secure:   true,
			Expires:  framework.Sessions[sessionId].ExpiredAt,
		})

		omittedUser := OmitUser(user)
		jsonResponse(w, omittedUser)
	}
}

func Logout(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// NOTE: cookie is checked in AuthMiddleware
		cookie, _ := r.Cookie("session")
		sessionId := cookie.Value
		delete(framework.Sessions, sessionId)
		jsonResponse(w, nil)
	}
}
