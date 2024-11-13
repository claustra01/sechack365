package handler

import (
	"encoding/json"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/openapi"
)

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
