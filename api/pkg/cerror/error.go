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

// image handler
var ErrInvalidFileType = fmt.Errorf("invalid file type")

// inbox handler
var ErrInvalidActivityType = fmt.Errorf("invalid activity type")
var ErrInvalidActivityObject = fmt.Errorf("invalid activity object")

// nostr service
var ErrNostrRelayResNotOk = fmt.Errorf("nostr relay response is not ok")

// bcrypt
var ErrGeneratePasswordHash = fmt.Errorf("failed to generate password hash")

// httpsig
var ErrUnknownKeyType = fmt.Errorf("unknown key type")
var ErrInvalidKeyPem = fmt.Errorf("invalid key pem")
var ErrInvalidHttpSig = fmt.Errorf("invalid http signature")

// schnorr
var ErrInvalidNostrEventId = fmt.Errorf("invalid nostr event id")
var ErrInvalidNostrEventSig = fmt.Errorf("invalid nostr event signature")

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}
