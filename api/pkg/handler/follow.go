package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/openapi"
	"github.com/claustra01/sechack365/pkg/util"

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
			// get keyId and privKey
			keyId := c.Controllers.ActivityPub.NewKeyIdUrl(user.Identifiers.Activitypub.Host, user.Id)
			privKeyPem, err := c.Controllers.User.GetActivityPubPrivKey(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			privKey, _, err := util.DecodePem(privKeyPem)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// resolve remote actor
			targetUrl, err := c.Controllers.ActivityPub.ResolveWebfinger(target.Identifiers.Activitypub.LocalUsername, target.Identifiers.Activitypub.Host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			targetActor, err := c.Controllers.ActivityPub.ResolveRemoteActor(targetUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// send activity
			follow, err := c.Controllers.Follow.FindFollowByFollowerAndTarget(user.Id, target.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			activity := &model.ApActivity{
				Context: *c.Controllers.ActivityPub.NewApContext(),
				Type:    model.ActivityTypeFollow,
				Id:      fmt.Sprintf("https://%s/follows/%s", user.Identifiers.Activitypub.Host, follow.Id),
				Actor:   c.Controllers.ActivityPub.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id),
				Object:  targetUrl,
			}
			if _, err := c.Controllers.ActivityPub.SendActivity(keyId, privKey, targetActor.Inbox, activity); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
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
			if err := c.Controllers.Nostr.PublishFollow(privKey, pubKeys); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
		}

		// success
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

		// check target user's protocol
		target, err := c.Controllers.User.FindById(unfollowRequestBody.TargetId)
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

		// delete follow
		if err := c.Controllers.Follow.Delete(user.Id, unfollowRequestBody.TargetId); err != nil {
			c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
			returnError(w, http.StatusInternalServerError)
			return
		}

		// activitypub remote unfollow
		if target.Protocol == model.ProtocolActivityPub {
			// get keyId and privKey
			keyId := c.Controllers.ActivityPub.NewKeyIdUrl(user.Identifiers.Activitypub.Host, user.Id)
			privKeyPem, err := c.Controllers.User.GetActivityPubPrivKey(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			privKey, _, err := util.DecodePem(privKeyPem)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// resolve remote actor
			targetUrl, err := c.Controllers.ActivityPub.ResolveWebfinger(target.Identifiers.Activitypub.LocalUsername, target.Identifiers.Activitypub.Host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			targetActor, err := c.Controllers.ActivityPub.ResolveRemoteActor(targetUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// send activity
			follow, err := c.Controllers.Follow.FindFollowByFollowerAndTarget(user.Id, target.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			activity := &model.ApActivity{
				Context: *c.Controllers.ActivityPub.NewApContext(),
				Type:    model.ActivityTypeUndo,
				Id:      fmt.Sprintf("https://%s/follows/%s", user.Identifiers.Activitypub.Host, follow.Id),
				Actor:   c.Controllers.ActivityPub.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id),
				Object: &model.ApActivity{
					Type:   model.ActivityTypeFollow,
					Id:     fmt.Sprintf("https://%s/follows/%s", user.Identifiers.Activitypub.Host, follow.Id),
					Actor:  c.Controllers.ActivityPub.NewActorUrl(user.Identifiers.Activitypub.Host, user.Id),
					Object: targetUrl,
				},
			}
			if _, err := c.Controllers.ActivityPub.SendActivity(keyId, privKey, targetActor.Inbox, activity); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to create follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
		}

		// nostr remote unfollow
		if target.Protocol == model.ProtocolNostr {
			pubKeys, err := c.Controllers.Follow.FindNostrFollowPublicKeys(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			privKey, err := c.Controllers.User.GetNostrPrivKey(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if err := c.Controllers.Nostr.PublishFollow(privKey, pubKeys); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to delete follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			// success
			returnResponse(w, http.StatusNoContent, ContentTypeJson, nil)
			return
		}

		returnResponse(w, http.StatusNoContent, ContentTypeJson, nil)
	}
}
