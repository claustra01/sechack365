package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/openapi"
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

		if postRequsetBody.Content == "" {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrEmptyContent, "failed to create post"))
			returnError(w, http.StatusBadRequest)
			return
		}

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

		if err := c.Controllers.Post.Create(user.Id, postRequsetBody.Content); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		privKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create post"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if err := c.Controllers.Nostr.PostText(privKey, postRequsetBody.Content); err != nil {
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
