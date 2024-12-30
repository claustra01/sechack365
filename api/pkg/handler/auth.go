package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

func Register(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authRequestBody openapi.Auth
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &authRequestBody)
		if err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusBadRequest)
			return
		}

		pattern := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
		if !pattern.MatchString(authRequestBody.Username) {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidUsername, "failed to register user"))
			returnError(w, http.StatusBadRequest)
			return
		}

		user, err := c.Controllers.User.FindByLocalUsername(authRequestBody.Username)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if user != nil {
			c.Logger.Warn("Conflict", "Error", cerror.Wrap(cerror.ErrUserAlreadyExists, "failed to register user"))
			returnError(w, http.StatusConflict)
			return
		}

		// TODO: set default icon url
		if err := c.Controllers.User.CreateLocalUser(authRequestBody.Username, authRequestBody.Password, authRequestBody.Username, "", "https://placehold.jp/150x150.png", c.Config.Host); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// create nostr profile
		user, err = c.Controllers.User.FindByLocalUsername(authRequestBody.Username)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		privKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		_, hexPrivKey, err := util.DecodeBech32(privKey)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		profile := &model.NostrProfile{
			Name:        user.Username,
			DisplayName: user.DisplayName,
			About:       user.Profile,
			Picture:     user.Icon,
		}
		if err := c.Controllers.Nostr.PostUserProfile(hexPrivKey, profile); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to register user"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		returnResponse(w, http.StatusCreated, ContentTypeJson, nil)
	}
}

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
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to login"))
			returnError(w, http.StatusBadRequest)
			return
		}

		user, err := c.Controllers.User.FindWithHashedPassword(authRequestBody.Username)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to login"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if user == nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(cerror.ErrInvalidPassword, "failed to login"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		if err := util.CompareHashAndPassword(user.HashedPassword, authRequestBody.Password); err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(cerror.ErrInvalidPassword, "failed to login"))
			returnError(w, http.StatusUnauthorized)
			return
		}

		for _, session := range framework.Sessions {
			if session.UserId == user.Id && session.ExpiredAt.Before(time.Now()) {
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
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			Expires:  framework.Sessions[sessionId].ExpiredAt,
		})

		returnResponse(w, http.StatusNoContent, ContentTypeJson, nil)
	}
}

func Logout(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// NOTE: cookie is checked in AuthMiddleware
		cookie, _ := r.Cookie("session")
		sessionId := cookie.Value
		delete(framework.Sessions, sessionId)

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
		})

		returnResponse(w, http.StatusNoContent, ContentTypeJson, nil)
	}
}
