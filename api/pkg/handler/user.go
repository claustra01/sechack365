package handler

import (
	"net/http"
	"regexp"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
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
		pattern := regexp.MustCompile(`^([a-zA-Z0-9_]+)@?([a-zA-Z0-9.]+)?$`)
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

		// save cache
		resolvedUser, err := c.Controllers.User.Insert(actor.PreferredUsername, "", host, actor.Name, actor.Summary)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if _, err := c.Controllers.ApUserIdentifier.Insert(resolvedUser.Id, actor.PublicKey.PublicKeyPem); err != nil {
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
