package handler

import (
	"net/http"
	"regexp"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
)

func NodeinfoLinks(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodeinfo := activitypub.BuildNodeInfoLinksSchema(c.Config.Host)
		jsonResponse(w, nodeinfo)
	}
}

func WebfingerLinks(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")

		pattern := regexp.MustCompile(`^acct:([a-zA-Z0-9_]+)@([a-zA-Z0-9.]+)$`)
		matches := pattern.FindStringSubmatch(resource)
		if len(matches) != 3 || matches[2] != c.Config.Host {
			returnBadRequest(w, c.Logger, cerror.ErrInvalidResourceQuery(resource))
			return
		}

		user, err := c.Controllers.User.FindByUsername(matches[1], c.Config.Host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if user == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}

		webfinger := activitypub.BuildWebfingerActorLinksSchema(user.Username, c.Config.Host)
		jsonCustomContentTypeResponse(w, webfinger, "application/jrd+json")
	}
}
