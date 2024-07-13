package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

func Nodeinfo(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Host:", "host", c.Config.Host)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"version":"0.1.0"}`)
	}
}

func Webfinger(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"subject":"acct:admin@localhost","links":[{"rel":"self","type":"application/activity+json","href":"http://localhost/api/v1/actor/admin"}]}`)
	}
}
