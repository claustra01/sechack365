package handler

import (
	"net/http"
	"regexp"
	"time"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

func OmitUser(user *model.User) *openapi.User {
	createdAt, err := util.StrToTime(user.CreatedAt)
	if err != nil {
		panic(err)
	}
	updatedAt, err := util.StrToTime(user.UpdatedAt)
	if err != nil {
		panic(err)
	}
	return &openapi.User{
		Id:          user.Id,
		Username:    user.Username,
		Host:        user.Host,
		Protocol:    user.Protocol,
		DisplayName: user.DisplayName,
		Profile:     user.Profile,
		Icon:        user.Icon,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func GetAllUsers(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		var omittedUsers []*openapi.User
		for _, user := range users {
			omittedUsers = append(omittedUsers, OmitUser(user))
		}
		jsonResponse(w, omittedUsers)
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
			omittedUser := OmitUser(user)
			jsonResponse(w, omittedUser)
		default:
			returnBadRequest(w, c.Logger, cerror.ErrInvalidAcceptHeader)
		}
	}
}

// FIXME: nostrユーザーの捜索は本来hashではなくnpub...の文字列で行うべき
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
			omittedUser := OmitUser(user)
			jsonResponse(w, omittedUser)
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
				omittedUser := OmitUser(user)
				jsonResponse(w, omittedUser)
				return
			}

			// activitypub
			// fetch from remote
			link, err := c.Controllers.ActivityPub.ResolveWebfinger(username, host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(link)
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
			omittedUser := OmitUser(user)
			jsonResponse(w, omittedUser)
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
			omittedUser := OmitUser(user)
			jsonResponse(w, omittedUser)
			return
		}

		// activitypub
		// fetch from remote
		link, err := c.Controllers.ActivityPub.ResolveWebfinger(username, host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(link)
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
		omittedUser := OmitUser(user)
		jsonResponse(w, omittedUser)
	}
}
