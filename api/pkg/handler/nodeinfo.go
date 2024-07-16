package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/framework"
)

func Nodeinfo2_0(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodeinfo := activitypub.GetNodeInfo()
		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(nodeinfo)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			panic(err)
		}
		fmt.Fprint(w, string(data))
	}
}
