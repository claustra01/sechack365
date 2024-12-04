package handler

import (
	"net/http"
	"regexp"
	"strings"
	"time"

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
		acceptHeader := r.Header.Get("Accept")

		user, err := c.Controllers.User.FindById(id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if user == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}
		switch acceptHeader {
		case "application/activity+json":
			actor := c.Controllers.ActivityPub.NewActor(*user)
			jsonCustomContentTypeResponse(w, actor, "application/activity+json")
		default:
			jsonResponse(w, user)
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
		jsonResponse(w, user)
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
		jsonResponse(w, follows)
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
		jsonResponse(w, followers)
	}
}

func GetUserPosts(c *framework.Context) http.HandlerFunc {
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
		// parse username
		usernameWithHost := r.PathValue("username")
		pattern := regexp.MustCompile(`^(@([a-zA-Z0-9_]+)(@[a-zA-Z0-9-.]+)?)|(npub[a-z0-9]+)$`)
		matches := pattern.FindStringSubmatch(usernameWithHost)

		// local/remote username
		username := matches[2]

		// remote host
		host := matches[3]
		if host != "" {
			host = strings.TrimLeft(host, "@")
		}

		// nostr public key
		npub := matches[4]

		// local user
		if (host == "" || host == c.Config.Host) && npub == "" {
			user, err := c.Controllers.User.FindByLocalUsername(username)
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

		// activitypub
		if host != "" {
			// check cache
			cachedUser, err := c.Controllers.User.FindByApUsername(username, host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if cachedUser != nil {
				cacheTime, err := util.StrToTime(cachedUser.UpdatedAt)
				if err == nil && util.CalcSubTime(time.Now(), cacheTime) < 24*time.Hour {
					jsonResponse(w, cachedUser)
					return
				}
			}

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

			u := &model.User{
				DisplayName: actor.Name,
				Profile:     actor.Summary,
				Icon:        actor.Icon.Url,
			}
			i := &model.ApUserIdentifier{
				LocalUsername: username,
				Host:          host,
			}

			// create user
			if cachedUser == nil {
				user, err := c.Controllers.User.CreateRemoteApUser(u, i)
				if err != nil {
					returnInternalServerError(w, c.Logger, err)
					return
				}
				jsonResponse(w, user)
				return
			}

			// update cache
			user, err := c.Controllers.User.UpdateRemoteApUser(u, i)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			jsonResponse(w, user)
			return
		}

		// nostr
		if npub != "" {
			// check cache
			cachedUser, err := c.Controllers.User.FindByNostrPublicKey(npub)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if cachedUser != nil {
				cacheTime, err := util.StrToTime(cachedUser.UpdatedAt)
				if err == nil && util.CalcSubTime(time.Now(), cacheTime) < 24*time.Hour {
					jsonResponse(w, cachedUser)
					return
				}
			}

			// decode bech32
			hrp, hexStr, err := util.DecodeBech32(npub)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if hrp != "npub" {
				returnBadRequest(w, c.Logger, cerror.ErrInvalidNostrKey)
				return
			}

			// fetch from remote
			profile, err := c.Controllers.Nostr.GetUserProfile(hexStr)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			if profile == nil {
				returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
				return
			}

			u := &model.User{
				DisplayName: profile.DisplayName,
				Profile:     profile.About,
				Icon:        profile.Picture,
			}
			i := &model.NostrUserIdentifier{
				PublicKey: npub,
			}

			// create user
			if cachedUser == nil {
				user, err := c.Controllers.User.CreateRemoteNostrUser(u, i)
				if err != nil {
					returnInternalServerError(w, c.Logger, err)
					return
				}
				jsonResponse(w, user)
				return
			}

			// update cache
			user, err := c.Controllers.User.UpdateRemoteNostrUser(u, i)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			jsonResponse(w, user)
			return
		}

		// return 404
		returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
	}
}
