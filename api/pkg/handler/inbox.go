package handler

import (
	"encoding/json"
	"net/http"
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
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(cerror.ErrInvalidHttpSig, "failed to verify http signature"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		keyId := match[1]
		actor, err := c.Controllers.ActivityPub.ResolveRemoteActor(keyId)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		pubKeyPem := actor.PublicKey.PublicKeyPem

		// verify signature
		_, pubKey, err := util.DecodePem(pubKeyPem)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature"))
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
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature"))
			// returnError(w, http.StatusUnauthorized)
			// return
		}

		// parse activity
		var activity map[string]interface{}
		if err := json.Unmarshal(body, &activity); err != nil {
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(err, "failed to parse activity"))
			returnError(w, http.StatusBadRequest)
			return
		}

		switch activity["type"] {
		case model.ActivityTypeFollow:
			// TODO: accept follow request
			returnError(w, http.StatusInternalServerError)
		case model.ActivityTypeAccept:
		case model.ActivityTypeReject:
		default:
			c.Logger.Warn("Bad Request", "Error", cerror.Wrap(cerror.ErrInvalidActivityType, "failed to parse activity"))
			returnError(w, http.StatusBadRequest)
			return
		}

		returnResponse(w, http.StatusAccepted, ContentTypeJson, nil)
	}
}
