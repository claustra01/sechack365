package handler

import (
	"net/http"
	"regexp"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
)

func NodeinfoLinks(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodeinfo := c.Controllers.Webfinger.NewNodeInfoLinks(c.Config.Host)
		returnResponse(w, http.StatusOK, ContentTypeJson, nodeinfo)
	}
}

func WebfingerLinks(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")

		pattern := regexp.MustCompile(`^acct:([a-zA-Z0-9_]+)@([a-zA-Z0-9-.]+)$`)
		matches := pattern.FindStringSubmatch(resource)
		if len(matches) != 3 || matches[2] != c.Config.Host {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidResourseQuery, "failed to resolve webfinger"))
			returnError(w, http.StatusBadRequest)
			return
		}

		user, err := c.Controllers.User.FindByLocalUsername(matches[1])
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to resolve webfinger"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if user == nil {
			c.Logger.Warn("Not Found", "Error", cerror.Wrap(cerror.ErrUserNotFound, "failed to resolve webfinger"))
			returnError(w, http.StatusNotFound)
			return
		}

		webfinger := c.Controllers.Webfinger.NewWebfingerActorLinks(c.Config.Host, user.Id, user.Username)
		returnResponse(w, http.StatusOK, ContentTypeJrdJson, webfinger)
	}
}
