package handler

import (
	"net/http"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
)

func GetActor(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("username")
		user, err := c.Controllers.ApUser.FindByUsername(name, c.Config.Host)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if user == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}
		actor := activitypub.BuildActorSchema(*user)
		jsonCustomContentTypeResponse(w, actor, "application/activity+json")
	}
}
