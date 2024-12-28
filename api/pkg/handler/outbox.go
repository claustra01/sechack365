package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func ActorOutbox(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.PathValue("username")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnError(w, http.StatusInternalServerError)
			return
		}
		log.Println(username, string(body))
		returnError(w, http.StatusInternalServerError)
	}
}
