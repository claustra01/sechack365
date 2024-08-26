package cerror

import "fmt"

var ErrUserNotFound = fmt.Errorf("user not found")

var ErrGeneratePasswordHash = fmt.Errorf("failed to generate password hash")
var ErrGenerateRsaKey = fmt.Errorf("failed to generate RSA key pair")
var ErrEncodePublicKey = fmt.Errorf("failed to encode public key")
var ErrEncodePrivateKey = fmt.Errorf("failed to encode private key")
var ErrDecodePublicKey = fmt.Errorf("failed to decode public key")
var ErrDecodePrivateKey = fmt.Errorf("failed to decode private key")
var ErrInvalidKeyPair = fmt.Errorf("invalid key pair")

func ErrInvalidResourceQuery(q string) error {
	return fmt.Errorf("invalid resource query: %s", q)
}
