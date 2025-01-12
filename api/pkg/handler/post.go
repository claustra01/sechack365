package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

// NOTE: timeline limit per request
const TimelineLimit = 10

func bindPost(p *model.PostWithUser) openapi.Post {
	post := openapi.Post{
		Id:      p.Id,
		Content: p.Content,
		User: openapi.SimpleUser{
			Username:    p.User.Username,
			Protocol:    p.User.Protocol,
			DisplayName: p.User.DisplayName,
			Icon:        p.User.Icon,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return post
}

func CreatePost(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postRequsetBody openapi.Newpost
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &postRequsetBody)
		if err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusBadRequest)
			return
		}

		// validate content
		if postRequsetBody.Content == "" || len(postRequsetBody.Content) > 200 {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrEmptyContent, "failed to create post"))
			returnError(w, http.StatusBadRequest)
			return
		}

		// get current user
		user, err := c.CurrentUser(r)
		if errors.Is(err, cerror.ErrUserNotFound) {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusUnauthorized)
			return
		} else if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// generate id
		uuid := util.NewUuid().String()

		// post to activitypub
		keyId := c.Controllers.ActivityPub.NewKeyIdUrl(user.Identifiers.Activitypub.Host, user.Id)
		privKeyPem, err := c.Controllers.User.GetActivityPubPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		apPrivKey, _, err := util.DecodePem(privKeyPem)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		followers, err := c.Controllers.Follow.FindActivityPubRemoteFollowers(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		for _, follower := range followers {
			f := strings.Split(follower, "@")
			targetUrl, err := c.Controllers.ActivityPub.ResolveWebfinger(f[0], f[1])
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			targetActor, err := c.Controllers.ActivityPub.ResolveRemoteActor(targetUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			activity := &model.ApActivity{
				Context: *c.Controllers.ActivityPub.NewApContext(),
				Type:    model.ActivityTypeCreate,
				Id:      fmt.Sprintf("https://%s/posts/%s", user.Identifiers.Activitypub.Host, uuid),
				Actor:   c.Controllers.ActivityPub.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id),
				Object: &model.ApNoteActivity{
					Id:        fmt.Sprintf("https://%s/posts/%s", user.Identifiers.Activitypub.Host, uuid),
					Type:      model.ActivityTypeNote,
					Content:   postRequsetBody.Content,
					Published: time.Now(),
				},
			}
			if _, err := c.Controllers.ActivityPub.SendActivity(keyId, apPrivKey, targetActor.Inbox, activity); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
				returnError(w, http.StatusInternalServerError)
				return
			}
		}

		// post to nostr
		nostrPrivKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if err := c.Controllers.Nostr.PublishPost(nostrPrivKey, postRequsetBody.Content); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// create
		if err := c.Controllers.Post.Create(uuid, user.Id, postRequsetBody.Content); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		returnResponse(w, http.StatusCreated, ContentTypeJson, nil)
	}
}

func GetPost(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidPathParam, "failed to get post"))
			returnError(w, http.StatusBadRequest)
			return
		}
		post, err := c.Controllers.Post.FindById(id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to get post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if post == nil {
			c.Logger.Warn("Not Found", "Error", cerror.Wrap(cerror.ErrPostNotFound, "failed to get post"))
			returnError(w, http.StatusNotFound)
			return
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, bindPost(post))
	}
}

func GetTimeline(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		offsetStr := r.URL.Query().Get("offset")
		offset := 0
		if offsetStr != "" {
			offset = int(offset)
		}
		posts, err := c.Controllers.Post.FindTimeline(offset, TimelineLimit)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to get timeline"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		var timeline []openapi.Post
		for _, p := range posts {
			timeline = append(timeline, bindPost(p))
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, timeline)
	}
}
