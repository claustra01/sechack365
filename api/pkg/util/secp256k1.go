package util

import (
	"encoding/hex"

	"github.com/decred/dcrd/dcrec/secp256k1"
)

func GenerateNostrKeyPair() (string, string, error) {
	rawPrivateKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}
	privateKey := hex.EncodeToString(rawPrivateKey.Serialize())

	rawPublicKey := rawPrivateKey.PubKey()
	// prefix (0x02) is removed in nostr public key
	publicKey := hex.EncodeToString(rawPublicKey.SerializeCompressed())[2:]

	return string(privateKey), string(publicKey), nil
}
