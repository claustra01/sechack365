package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/claustra01/sechack365/pkg/activitypub"
	"github.com/claustra01/sechack365/pkg/util"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
)

type FollowRequestBody struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

func CreateFollow(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var followRequestBody FollowRequestBody
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &followRequestBody)
		if err != nil {
			returnBadRequest(w, c.Logger, err)
			return
		}

		// check follower protocol
		follower, err := c.Controllers.User.FindById(followRequestBody.Follower)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if follower == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}
		if follower.Protocol != model.ProtocolLocal {
			returnBadRequest(w, c.Logger, cerror.ErrInvalidFollowRequest)
			return
		}

		// check followee protocol
		followee, err := c.Controllers.User.FindById(followRequestBody.Followee)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}
		if followee == nil {
			returnNotFound(w, c.Logger, cerror.ErrUserNotFound)
			return
		}
		if followee.Protocol != model.ProtocolLocal && followee.Protocol != model.ProtocolActivityPub {
			returnBadRequest(w, c.Logger, cerror.ErrInvalidFollowRequest)
			return
		}

		// create follow
		follow, err := c.Controllers.Follow.Create(followRequestBody.Follower, followRequestBody.Followee)
		if err != nil {
			returnInternalServerError(w, c.Logger, err)
			return
		}

		switch followee.Protocol {
		case model.ProtocolLocal:
			// nop
		case model.ProtocolActivityPub:
			// get signer key params
			keyId := activitypub.BuildKeyIdUrl(follower.Host, follower.Username)
			followerIdentifier, err := c.Controllers.ApUserIdentifier.FindById(follower.Id)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			privateKey, err := util.DecodePrivateKeyPem(followerIdentifier.PrivateKey)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}

			// get followee actor url
			followeeUrl, err := activitypub.ResolveWebfinger(followee.Username, followee.Host)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}

			// send follow activity
			followActivity := activitypub.BuildFollowActivitySchema(follow.Id, follower.Host, follower.Id, followeeUrl)
			followeeActor, err := activitypub.ResolveRemoteActor(followActivity.Object)
			if err != nil {
				returnInternalServerError(w, c.Logger, err)
				return
			}
			signParams := activitypub.SignParams{
				Host:       c.Config.Host,
				KeyId:      keyId,
				PrivateKey: privateKey,
			}
			respBody, err := activitypub.SendActivity(followeeActor.Inbox, followActivity, signParams)
			if err != nil {
				c.Logger.Error("Remote follow error", "ERROR", string(respBody))
				returnInternalServerError(w, c.Logger, err)
				return
			}

			// FIXME: This is debug
			log.Println(string(respBody))
			return
		}
	}
}
