package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/framework"
)

func Nodeinfo(c *framework.Context) http.HandlerFunc {
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

func Webfinger(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")

		// mock actor
		exceptedResource := fmt.Sprintf("acct:%s@%s", "mock", c.Config.Host)
		if resource != exceptedResource {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		webfinger := activitypub.GetWebfingerActor("mock", c.Config.Host)
		data, err := json.Marshal(webfinger)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(data))
	}
}
