package handler

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/framework"
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
		fmt.Println(pubKeyPem)

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
		keyname, err := util.HttpSigVerify(r, body, pubKey)
		if err != nil {
			c.Logger.Warn("Unauthorized", "Error", cerror.Wrap(err, "failed to verify http signature"))
			returnError(w, http.StatusUnauthorized)
			return
		}
		fmt.Println(keyname)

		// debug
		returnError(w, http.StatusInternalServerError)
	}
}
