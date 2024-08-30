package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func ActorInbox(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		log.Println(username, string(body))
		returnInternalServerError(w, c.Logger, nil)
	}
}
