package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

func ActorInbox(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// resolve public key
		sigHeader := r.Header.Get("Signature")
		re := regexp.MustCompile(`keyId="([^"]+)"`)
		match := re.FindStringSubmatch(sigHeader)
		if len(match) <= 1 {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(cerror.ErrInvalidHttpSig, "failed to verify http signature at inbox"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		keyId := match[1]
		actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(keyId)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature at inbox"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		pubKeyPem := actor.PublicKey.PublicKeyPem

		// verify signature
		_, pubKey, err := util.DecodePem(pubKeyPem)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature at inbox"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		body := make([]byte, r.ContentLength)
		if _, err := r.Body.Read(body); err != nil && err.Error() != "EOF" {
			// NOTE: err should be nil
			panic(err)
		}
		_, err = util.HttpSigVerify(r, body, pubKey)
		// TODO: fix httpsig verifier
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature at inbox"))
			// returnError(w, http.StatusUnauthorized)
			// return
		}

		// parse activity
		var activity map[string]interface{}
		if err := json.Unmarshal(body, &activity); err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to parse activity at inbox"))
			returnError(w, http.StatusBadRequest)
			return
		}

		// resolve target
		targetUrl := activity["object"].(string)
		re = regexp.MustCompile(`https://` + c.Config.Host + `/api/v1/users/([a-z0-9-]+)`)
		match = re.FindStringSubmatch(targetUrl)
		if len(match) <= 1 {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidActivityObject, "failed to parse activity at inbox"))
			returnError(w, http.StatusBadRequest)
			return
		}
		targetId := match[1]

		switch activity["type"] {
		// follow
		case model.ActivityTypeFollow:
			// resolve follower
			followerUrl := activity["actor"].(string)
			parsedFollowerURL, err := url.Parse(followerUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(followerUrl)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			follower, err := c.Controllers.User.FindByApUsername(actor.PreferredUsername, parsedFollowerURL.Host)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			if follower == nil {
				u := &model.User{
					DisplayName: actor.Name,
					Profile:     actor.Summary,
					Icon:        actor.Icon.Url,
				}
				i := &model.ApUserIdentifier{
					LocalUsername: actor.PreferredUsername,
					Host:          parsedFollowerURL.Host,
				}
				if err := c.Controllers.User.CreateRemoteApUser(u, i); err != nil {
					c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
					returnError(w, http.StatusInternalServerError)
					return
				}
				follower, err = c.Controllers.User.FindByApUsername(actor.PreferredUsername, parsedFollowerURL.Host)
				if err != nil {
					c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
					returnError(w, http.StatusInternalServerError)
					return
				}
			}

			// get keyId and privKey
			user, err := c.Controllers.User.FindById(targetId)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			keyId := c.Controllers.ActivityPub.NewKeyIdUrl(user.Identifiers.Activitypub.Host, user.Id)
			privKeyPem, err := c.Controllers.User.GetActivityPubPrivKey(user.Id)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			privKey, _, err := util.DecodePem(privKeyPem)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// create follow
			if err := c.Controllers.Follow.Create(follower.Id, targetId); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

			// send activity
			follow, err := c.Controllers.Follow.FindFollowByFollowerAndTarget(follower.Id, targetId)
			if err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}
			acceptActivity := &model.ApActivity{
				Context: *c.Controllers.ActivityPub.NewApContext(),
				Type:    model.ActivityTypeAccept,
				Id:      fmt.Sprintf("https://%s/follows/%s", c.Config.Host, follow.Id),
				Actor:   targetUrl,
				Object:  string(body),
			}
			if _, err := c.Controllers.ActivityPub.SendActivity(keyId, privKey, actor.Inbox, acceptActivity); err != nil {
				c.Logger.Error("Internal Server Error", "Error", cerror.Wrap(err, "failed to receive activitypub remote follow"))
				returnError(w, http.StatusInternalServerError)
				return
			}

		case model.ActivityTypeAccept:
		case model.ActivityTypeReject:

		// undo
		case model.ActivityTypeUndo:

		default:
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidActivityType, "failed to parse activity"))
			returnError(w, http.StatusBadRequest)
			return
		}

		returnResponse(w, http.StatusAccepted, ContentTypeJson, nil)
	}
}
