package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

// TODO: save session to database or cache
type Session struct {
	Id        util.Uuid
	UserId    string
	CreatedAt time.Time
	ExpiredAt time.Time
}

var sessions = make(map[util.Uuid]Session)

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

		sessionId := util.NewUuid()
		sessions[sessionId] = Session{
			Id:        sessionId,
			UserId:    user.Id,
			CreatedAt: time.Now(),
			ExpiredAt: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    sessionId.String(),
			HttpOnly: true,
			Secure:   true,
			Expires:  sessions[sessionId].ExpiredAt,
		})

		omittedUser := OmitUser(user)
		jsonResponse(w, omittedUser)
	}
}
