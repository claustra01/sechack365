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
		header := r.Header
		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		log.Println(username)
		log.Println(header.Get("Signature"))
		log.Println(string(body))
		returnInternalServerError(w, c.Logger, nil)
	}
}
