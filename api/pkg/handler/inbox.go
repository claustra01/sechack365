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
		headers := r.Header
		digest := headers.Get("Digest")
		signature := headers.Get("Signature")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		// FIXME: remove debug log
		log.Println(username)
		log.Println(digest)
		log.Println(signature)
		log.Println(string(body))
		returnInternalServerError(w, c.Logger, nil)
	}
}
