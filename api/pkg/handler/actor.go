package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/framework"
)

func MockActor(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actor, err := activitypub.GetActor("mock", c.Config.Host)
		if err != nil {
			// TODO: switch error (404, 500, etc...)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(actor)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/activity+json")
		fmt.Fprint(w, string(data))
	}
}
