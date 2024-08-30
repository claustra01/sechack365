package handler

import (
	"net/http"
	"regexp"
	"time"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

func GetAllUsers(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		jsonResponse(w, users)
	}
}

func LookupUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usernameWithHost := r.PathValue("username")
		pattern := regexp.MustCompile(`^([a-zA-Z0-9_]+)@?([a-zA-Z0-9-.]+)?$`)
		matches := pattern.FindStringSubmatch(usernameWithHost)

		if len(matches) != 3 {
			returnBadRequest(w, c.Logger, cerror.ErrInvalidResourceQuery(usernameWithHost))
			return
		}

		username := matches[1]
		host := matches[2]

		// local user
		if host == "" || host == c.Config.Host {
			username := matches[1]
			user, err := c.Controllers.User.FindByUsername(username, c.Config.Host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if user == nil {
				returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
				return
			}
			jsonResponse(w, user)
			return
		}

		// check cache
		cachedUser, err := c.Controllers.User.FindByUsername(username, host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		if cachedUser != nil {
			// return cached user
			cachedTime, err := util.StrToTime(cachedUser.UpdatedAt)
			if err == nil && util.CalcSubTime(time.Now(), cachedTime) < 24*time.Hour {
				jsonResponse(w, cachedUser)
				return
			}

			// fetch from remote
			link, err := activitypub.ResolveWebfinger(username, host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			actor, err := activitypub.ResolveRemoteActor(link)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}

			// update cache
			if _, err := c.Controllers.User.UpdateRemoteUser(username, host, actor.Name, actor.Summary, actor.Icon.Url); err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			user, err := c.Controllers.User.FindByUsername(username, host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			jsonResponse(w, user)
		}

		// fetch from remote
		link, err := activitypub.ResolveWebfinger(username, host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		actor, err := activitypub.ResolveRemoteActor(link)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		// save cache
		if err := c.Controllers.Transaction.Begin(); err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		defer func() {
			if err := c.Controllers.Transaction.Rollback(); err != nil {
				c.Logger.Error("Transaction Rollback Failed", err)
			}
		}()
		resolvedUser, err := c.Controllers.User.Insert(actor.PreferredUsername, "", host, model.ProtocolActivityPub, actor.Name, actor.Summary, actor.Icon.Url)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if _, err := c.Controllers.ApUserIdentifier.Insert(resolvedUser.Id, actor.Id, actor.Inbox, actor.Outbox, actor.PublicKey.PublicKeyPem); err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if err := c.Controllers.Transaction.Commit(); err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		// return
		user, err := c.Controllers.User.FindByUsername(username, host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		jsonResponse(w, user)
	}
}
