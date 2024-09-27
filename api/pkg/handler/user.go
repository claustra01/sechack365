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

func GetUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		header := r.Header.Get("Accept")

		user, err := c.Controllers.User.FindById(id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if user == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}
		identifier, err := c.Controllers.ApUserIdentifier.FindById(user.Id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		switch header {
		case "application/activity+json":
			actor := activitypub.BuildActorSchema(*user, *identifier)
			jsonCustomContentTypeResponse(w, actor, "application/activity+json")
		case "application/json":
			jsonResponse(w, user)
		default:
			returnBadRequest(w, c.Logger, cerror.ErrInvalidAcceptHeader)
		}
	}
}

func LookupUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usernameWithHost := r.PathValue("username")
		pattern := regexp.MustCompile(`^([a-zA-Z0-9_]+)@?([a-zA-Z0-9-.]+)?$`)
		matches := pattern.FindStringSubmatch(usernameWithHost)

		if len(matches) != 3 {
			returnBadRequest(w, c.Logger, cerror.Wrap(cerror.ErrInvalidResourseQuery, usernameWithHost))
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

			// nostr
			if host == "nostr" {
				// fetch from remote
				profile, err := c.Controllers.Nostr.GetUserProfile(username)
				if err != nil {
					returnInternalServerError(w, c.Logger, err)
					return
				}
				if profile == nil {
					returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
					return
				}

				// update cache
				if _, err := c.Controllers.User.UpdateRemoteUser(username, host, profile.DisplayName, profile.About, profile.Picture); err != nil {
					returnInternalServerError(w, c.Logger, err)
					return
				}
				user, err := c.Controllers.User.FindByUsername(username, host)
				if err != nil {
					returnInternalServerError(w, c.Logger, err)
					return
				}
				jsonResponse(w, user)
				return
			}

			// activitypub
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

		// nostr
		if host == "nostr" {
			// fetch from remote
			profile, err := c.Controllers.Nostr.GetUserProfile(username)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if profile == nil {
				returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
				return
			}

			// save cache
			user, err := c.Controllers.User.CreateRemoteUser(username, host, model.ProtocolNostr, profile.DisplayName, profile.About, profile.Picture)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			jsonResponse(w, user)
			return
		}

		// activitypub
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
		user, err := c.Controllers.User.CreateRemoteUser(actor.PreferredUsername, host, model.ProtocolActivityPub, actor.Name, actor.Summary, actor.Icon.Url)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		jsonResponse(w, user)
	}
}
