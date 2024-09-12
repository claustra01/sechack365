package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrUserNotFound = fmt.Errorf("user not found")

var ErrInvalidResourseQuery = fmt.Errorf("invalid resource query")
var ErrResolveWebfinger = fmt.Errorf("failed to resolve webfinger")
var ErrResolveRemoteActor = fmt.Errorf("failed to resolve remote actor")

var ErrInvalidFollowRequest = fmt.Errorf("invalid follow request")
var ErrPushActivity = fmt.Errorf("failed to push activity")

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
