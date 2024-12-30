package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrInvalidQueryParam = fmt.Errorf("invalid query parameter")
var ErrInvalidPathParam = fmt.Errorf("invalid path parameter")

// auth handler
var ErrInvalidPassword = fmt.Errorf("invalid username or password")
var ErrUserAlreadyExists = fmt.Errorf("user already exists")
var ErrInvalidUsername = fmt.Errorf("invalid username")

// user handler
var ErrUserNotFound = fmt.Errorf("user not found")
var ErrResolveWebfinger = fmt.Errorf("failed to resolve webfinger")
var ErrResolveRemoteActor = fmt.Errorf("failed to resolve remote actor")
var ErrInvalidAcceptHeader = fmt.Errorf("invalid accept header")
var ErrInvalidNostrKey = fmt.Errorf("invalid nostr key")

var ErrInvalidFollowRequest = fmt.Errorf("invalid follow request")
var ErrPushActivity = fmt.Errorf("failed to push activity")

// post handler
var ErrEmptyContent = fmt.Errorf("empty content")
var ErrPostNotFound = fmt.Errorf("post not found")

// nostr service
var ErrNostrRelayResNotOk = fmt.Errorf("nostr relay response is not ok")

// utils
var ErrGeneratePasswordHash = fmt.Errorf("failed to generate password hash")
var ErrGenerateRsaKey = fmt.Errorf("failed to generate RSA key pair")
var ErrEncodePublicKey = fmt.Errorf("failed to encode public key")
var ErrEncodePrivateKey = fmt.Errorf("failed to encode private key")
var ErrDecodePublicKey = fmt.Errorf("failed to decode public key")
var ErrDecodePrivateKey = fmt.Errorf("failed to decode private key")
var ErrInvalidKeyPair = fmt.Errorf("invalid key pair")
var ErrGenerateSignature = fmt.Errorf("failed to generate signature")

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}
