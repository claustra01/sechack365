package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/decred/dcrd/dcrec/secp256k1/v3"
	"github.com/decred/dcrd/dcrec/secp256k1/v3/schnorr"
)

func GenerateNostrKeyPair() (string, string, error) {
	rawPrivKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return "", "", err
	}
	privKey := hex.EncodeToString(rawPrivKey.Serialize())

	rawPubKey := rawPrivKey.PubKey()
	// prefix (0x02) is removed in nostr public key
	pubKey := hex.EncodeToString(rawPubKey.SerializeCompressed())[2:]

	return string(privKey), string(pubKey), nil
}

func NostrVerify(pubKey string, digest string, sig string) (bool, error) {
	rawPubKey, err := hex.DecodeString("02" + pubKey)
	if err != nil {
		return false, err
	}
	pubKeyObj, err := secp256k1.ParsePubKey(rawPubKey)
	if err != nil {
		return false, err
	}
	rawSig, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}
	sigObj, err := schnorr.ParseSignature(rawSig)
	if err != nil {
		return false, err
	}
	hash, err := hex.DecodeString(digest)
	if err != nil {
		return false, err
	}
	return sigObj.Verify(hash, pubKeyObj), nil
}

func NostrSign(privKey string, pubKey string, createdAt time.Time, kind int, tags []model.NostrEventTag, content any) (*model.NostrEvent, error) {
	// serialize content
	rawContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	// calculate digest
	obj := []any{
		0,
		pubKey,
		createdAt.Unix(),
		kind,
		tags,
		content,
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(objStr)

	// signature
	rawPrivKey, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	sig, err := schnorr.Sign(secp256k1.PrivKeyFromBytes(rawPrivKey), hash[:])
	if err != nil {
		return nil, err
	}

	// return
	return &model.NostrEvent{
		Id:        hex.EncodeToString(hash[:]),
		Pubkey:    pubKey,
		CreatedAt: int(createdAt.Unix()),
		Kind:      kind,
		Tags:      tags,
		Content:   string(rawContent),
		Sig:       hex.EncodeToString(sig.Serialize()),
	}, nil
}
