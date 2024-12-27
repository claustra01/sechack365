package handler

import (
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
)

func Nodeinfo2_0(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := c.Controllers.User.FindAll()
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to get nodeinfo"))
			returnError(w, http.StatusInternalServerError)
		}
		nodeinfo := c.Controllers.ActivityPub.NewNodeInfo(len(users))
		returnResponse(w, http.StatusOK, ContentTypeJson, nodeinfo)
	}
}
