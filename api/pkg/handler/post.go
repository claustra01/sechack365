package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"
)

// NOTE: timeline limit per request
const TimelineLimit = 10

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
			returnBadRequest(w, c.Logger, err)
			return
		}

		if postRequsetBody.Text == "" {
			returnBadRequest(w, c.Logger, nil)
			return
		}

		user, err := c.CurrentUser(r)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		post, err := c.Controllers.Post.Create(user.Id, postRequsetBody.Text)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		jsonResponse(w, post)
	}
}

func GetPost(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		post, err := c.Controllers.Post.FindById(id)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		jsonResponse(w, post)
	}
}

func GetTimeline(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		log.Println(createdAt)

		posts, err := c.Controllers.Post.FindTimeline(createdAt, TimelineLimit)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		jsonResponse(w, posts)
	}
}
