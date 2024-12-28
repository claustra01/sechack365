package handler

import (
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

func bindUser(u *model.UserWithIdentifiers) openapi.User {
	var identifiers openapi.Identifiers
	if u.Identifiers.Activitypub != nil {
		identifiers.Activitypub = &openapi.ApIdentifier{
			LocalUsername: u.Identifiers.Activitypub.LocalUsername,
			Host:          u.Identifiers.Activitypub.Host,
			PublicKey:     u.Identifiers.Activitypub.PublicKey,
		}
	}
	if u.Identifiers.Nostr != nil {
		identifiers.Nostr = &openapi.NostrIdentifier{
			PublicKey: u.Identifiers.Nostr.PublicKey,
		}
	}
	user := openapi.User{
		Id:            u.Id,
		Username:      u.Username,
		Protocol:      u.Protocol,
		DisplayName:   u.DisplayName,
		Profile:       u.Profile,
		Icon:          u.Icon,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		Identifiers:   identifiers,
		PostCount:     u.PostCount,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}
	return user
}

func bindSimpleUser(u *model.User) openapi.SimpleUser {
	user := openapi.SimpleUser{
		Username:    u.Username,
		Protocol:    u.Protocol,
		DisplayName: u.DisplayName,
		Icon:        u.Icon,
	}
	return user
}

func GetAllUsers(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			c.Logger.Error("failed to find all users", "error", err)
			returnError(w, http.StatusInternalServerError)
			return
		}

		var usersResponse []openapi.User
		for _, u := range users {
			usersResponse = append(usersResponse, bindUser(u))
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, usersResponse)
	}
}

func GetUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		acceptHeader := r.Header.Get("Accept")

		if id == "" {
			c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidQueryParam, "failed to find user"))
			returnError(w, http.StatusBadRequest)
		}

		user, err := c.Controllers.User.FindById(id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to find user"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		if user == nil {
			c.Logger.Warn("Not Found", "error", cerror.Wrap(cerror.ErrUserNotFound, "failed to find user"))
			returnError(w, http.StatusNotFound)
			return
		}

		switch acceptHeader {
		case "application/activity+json":
			actor := c.Controllers.ActivityPub.NewActor(*user)
			returnResponse(w, http.StatusOK, ContentTypeApJson, actor)
		default:
			returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
		}
	}
}

func GetCurrentUser(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := c.CurrentUser(r)
		if errors.Is(err, cerror.ErrUserNotFound) {
			c.Logger.Warn("Not Found", "error", cerror.Wrap(err, "failed to find current user"))
			returnError(w, http.StatusNotFound)
			return
		} else if err != nil {
			c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to find current user"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
	}
}

func GetUserFollows(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidPathParam, "failed to find user follows"))
			returnError(w, http.StatusBadRequest)
			return
		}
		follows, err := c.Controllers.Follow.FindFollowsByUserId(id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to find user follows"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		var followsResponse []openapi.SimpleUser
		for _, f := range follows {
			followsResponse = append(followsResponse, bindSimpleUser(f))
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, followsResponse)
	}
}

func GetUserFollowers(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidPathParam, "failed to find user followers"))
			returnError(w, http.StatusBadRequest)
			return
		}
		followers, err := c.Controllers.Follow.FindFollowersByUserId(id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to find user followers"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		var followersResponse []openapi.SimpleUser
		for _, f := range followers {
			followersResponse = append(followersResponse, bindSimpleUser(f))
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, followersResponse)
	}
}

func GetUserPosts(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.PathValue("id")
		if userId == "" {
			c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidPathParam, "failed to find user posts"))
			returnError(w, http.StatusBadRequest)
			return
		}

		offsetStr := r.URL.Query().Get("offset")
		offset := 0
		if offsetStr != "" {
			offset = int(offset)
		}

		posts, err := c.Controllers.Post.FindUserTimeline(userId, offset, TimelineLimit)
		if err != nil {
			c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to find user posts"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		var postsResponse []openapi.Post
		for _, p := range posts {
			postsResponse = append(postsResponse, bindPost(p))
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, postsResponse)
	}
}

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
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if user == nil {
				c.Logger.Warn("Not Found", "error", cerror.Wrap(cerror.ErrUserNotFound, "failed to lookup user"))
				returnError(w, http.StatusNotFound)
				return
			}
			returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
			return
		}

		// activitypub
		if host != "" {
			// check cache
			cachedUser, err := c.Controllers.User.FindByApUsername(username, host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if cachedUser != nil {
				if err == nil && util.CalcSubTime(time.Now(), cachedUser.UpdatedAt) < 24*time.Hour {
					returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(cachedUser))
					return
				}
			}

			// fetch from remote
			link, err := c.Controllers.ActivityPub.ResolveWebfinger(username, host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(link)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
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
				if err := c.Controllers.User.CreateRemoteApUser(u, i); err != nil {
					c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
					returnError(w, http.StatusInternalServerError)
					return
				}
				user, err := c.Controllers.User.FindByApUsername(username, host)
				if err != nil {
					c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
					returnError(w, http.StatusInternalServerError)
					return
				}
				returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
				return
			}

			// update cache
			if err := c.Controllers.User.UpdateRemoteApUser(u, i); err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			user, err := c.Controllers.User.FindByApUsername(username, host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
			return
		}

		// nostr
		if npub != "" {
			// check cache
			cachedUser, err := c.Controllers.User.FindByNostrPublicKey(npub)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if cachedUser != nil {
				if err == nil && util.CalcSubTime(time.Now(), cachedUser.UpdatedAt) < 24*time.Hour {
					returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(cachedUser))
					return
				}
			}

			// decode bech32
			hrp, hexStr, err := util.DecodeBech32(npub)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if hrp != "npub" {
				c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidNostrKey, "failed to lookup user"))
				returnError(w, http.StatusBadRequest)
				return
			}

			// fetch from remote
			profile, err := c.Controllers.Nostr.GetUserProfile(hexStr)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if profile == nil {
				c.Logger.Warn("Not Found", "error", cerror.Wrap(cerror.ErrUserNotFound, "failed to lookup user"))
				returnError(w, http.StatusNotFound)
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
				if err := c.Controllers.User.CreateRemoteNostrUser(u, i); err != nil {
					c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
					returnError(w, http.StatusInternalServerError)
					return
				}
				user, err := c.Controllers.User.FindByNostrPublicKey(npub)
				if err != nil {
					c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
					returnError(w, http.StatusInternalServerError)
					return
				}
				returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
				return
			}

			// update cache
			if err := c.Controllers.User.UpdateRemoteNostrUser(u, i); err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			user, err := c.Controllers.User.FindByNostrPublicKey(npub)
			if err != nil {
				c.Logger.Error("Internal Server Error", "error", cerror.Wrap(err, "failed to lookup user"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			returnResponse(w, http.StatusOK, ContentTypeJson, bindUser(user))
			return
		}

		// return 404
		c.Logger.Warn("Not Found", "error", cerror.Wrap(cerror.ErrUserNotFound, "failed to lookup user"))
		returnError(w, http.StatusNotFound)
	}
}
