package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/claustra01/sechack365/pkg/openapi"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
)

func CreateFollow(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var followRequestBody openapi.Newfollow
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &followRequestBody)
		if err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusBadRequest)
			return
		}

		user, err := c.CurrentUser(r)
		if errors.Is(err, cerror.ErrUserNotFound) {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusUnauthorized)
			return
		} else if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// check target user's protocol
		target, err := c.Controllers.User.FindById(followRequestBody.TargetId)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}
		if target == nil {
			c.Logger.Warn("Not Found", "Error", cerror.Wrap(cerror.ErrUserNotFound, "failed to create follow"))
			returnError(w, http.StatusNotFound)
			return
		}

		// create follow
		if err := c.Controllers.Follow.Create(user.Id, target.Id); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// activitypub remote follow
		if target.Protocol == model.ProtocolActivityPub {
			// get signer key params
			// keyId := c.Controllers.ActivityPub.NewKeyIdUrl(follower.Host, follower.Username)
			// followerIdentifier, err := c.Controllers.ApUserIdentifier.FindById(follower.Id)
			// if err != nil {
			// 	returnInternalServerError(w, c.Logger, err)
			// 	return
			// }
			// privateKey, err := util.DecodePrivateKeyPem(followerIdentifier.PrivateKey)
			// if err != nil {
			// 	returnInternalServerError(w, c.Logger, err)
			// 	return
			// }

			// // get followee actor url
			// followeeUrl, err := c.Controllers.ActivityPub.ResolveWebfinger(followee.Username, followee.Host)
			// if err != nil {
			// 	returnInternalServerError(w, c.Logger, err)
			// 	return
			// }

			// // send follow activity
			// followActivity := c.Controllers.ActivityPub.NewFollowActivity(follow.Id, follower.Host, follower.Id, followeeUrl)
			// followeeActor, err := c.Controllers.ActivityPub.ResolveRemoteActor(followActivity.Object)
			// if err != nil {
			// 	returnInternalServerError(w, c.Logger, err)
			// 	return
			// }
			// respBody, err := c.Controllers.ActivityPub.SendActivity(followeeActor.Inbox, followActivity, c.Config.Host, keyId, privateKey)
			// if err != nil {
			// 	c.Logger.Error("Remote follow error", "ERROR", string(respBody))
			// 	returnInternalServerError(w, c.Logger, err)
			// 	return
			// }

			// // FIXME: This is debug
			// log.Println(string(respBody))
			// return
		}

		// nostr remote follow
		if target.Protocol == model.ProtocolNostr {
			pubKeys, err := c.Controllers.Follow.FindNostrFollowPublicKeys(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			privKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if err := c.Controllers.Nostr.PostFollow(privKey, pubKeys); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			// success
			returnResponse(w, http.StatusCreated, ContentTypeJson, nil)
			return
		}

		// local follow
		returnResponse(w, http.StatusCreated, ContentTypeJson, nil)
	}
}

func CheckIsFollowing(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		targetId := r.PathValue("id")
		if targetId == "" {
			c.Logger.Warn("Bad Request", "error", cerror.Wrap(cerror.ErrInvalidQueryParam, "failed to find user"))
			returnError(w, http.StatusBadRequest)
		}

		user, err := c.CurrentUser(r)
		if errors.Is(err, cerror.ErrUserNotFound) {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to check is following"))
			returnError(w, http.StatusUnauthorized)
			return
		} else if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to check is following"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		isFollowing, err := c.Controllers.Follow.CheckIsFollowing(user.Id, targetId)
		if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to check is following"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		resp := openapi.Found{
			Found: isFollowing,
		}
		returnResponse(w, http.StatusOK, ContentTypeJson, resp)
	}
}

func DeleteFollow(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		var unfollowRequestBody openapi.Newfollow
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		err := json.Unmarshal(body, &unfollowRequestBody)
		if err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to delete follow"))
			returnError(w, http.StatusBadRequest)
			return
		}

		user, err := c.CurrentUser(r)
		if errors.Is(err, cerror.ErrUserNotFound) {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to delete follow"))
			returnError(w, http.StatusUnauthorized)
			return
		} else if err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		if err := c.Controllers.Follow.Delete(user.Id, unfollowRequestBody.TargetId); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		returnResponse(w, http.StatusNoContent, ContentTypeJson, nil)
	}
}
