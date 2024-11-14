package handler

import (
	"log"
	"net/http"
	"regexp"
	"time"

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
			actor := c.Controllers.ActivityPub.NewActor(*user, *identifier)
			jsonCustomContentTypeResponse(w, actor, "application/activity+json")
		// case "application/json":
		// 	omittedUser := OmitUser(user)
		// 	jsonResponse(w, omittedUser)
		default:
			omittedUser := OmitUser(user)
			jsonResponse(w, omittedUser)
		}
	}
}

func GetCurrentUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.CurrentUser(r)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		omittedUser := OmitUser(user)
		jsonResponse(w, omittedUser)
	}
}

func GetUserFollows(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		follows, err := c.Controllers.Follow.FindFollowsByUserId(id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		var omittedFollows []*openapi.User
		for _, user := range follows {
			omittedFollows = append(omittedFollows, OmitUser(user))
		}
		log.Println(omittedFollows)
		jsonResponse(w, omittedFollows)
	}
}

func GetUserFollowers(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		followers, err := c.Controllers.Follow.FindFollowersByUserId(id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		var omittedFollowers []*openapi.User
		for _, user := range followers {
			omittedFollowers = append(omittedFollowers, OmitUser(user))
		}
		log.Println(omittedFollowers)
		jsonResponse(w, omittedFollowers)
	}
}

func GetUserTimeline(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")
		createdAtStr := r.URL.Query().Get("created_at")
		var createdAt time.Time
		if createdAtStr != "" {
			var err error
			createdAt, err = util.StrToTime(createdAtStr)
			if err != nil {
				returnBadRequest(w, c.Logger, err)
				return
			}
		} else {
			createdAt = time.Now()
		}

		posts, err := c.Controllers.Post.FindUserTimeline(userId, createdAt, TimelineLimit)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		jsonResponse(w, posts)
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
