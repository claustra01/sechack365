package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
)

func GeneratePrivateKey() string {
	params := btcec.S256().Params()
	one := new(big.Int).SetInt64(1)

	b := make([]byte, params.BitSize/8+8)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)

	return fmt.Sprintf("%064x", k.Bytes())
}

func GetPublicKey(sk string) (string, error) {
	b, err := hex.DecodeString(sk)
	if err != nil {
		return "", err
	}
	_, pk := btcec.PrivKeyFromBytes(b)
	return hex.EncodeToString(schnorr.SerializePubKey(pk)), nil
}

func GenerateNostrKeyPair() (string, string, error) {
	privKey := GeneratePrivateKey()
	pubKey, err := GetPublicKey(privKey)
	if err != nil {
		return "", "", cerror.Wrap(err, "failed to generate nostr key pair")
	}
	return privKey, pubKey, nil
}

func NostrSign(privKey string, createdAt time.Time, kind int, tags []model.NostrEventTag, content any) (*model.NostrEvent, error) {
	var contentStr string
	switch kind {
	// text note
	case 1:
		contentStr = content.(string)
	// others (e.g. profile)
	default:
		rawContent, err := json.Marshal(content)
		if err != nil {
			return nil, cerror.Wrap(err, "failed to sign nostr event")
		}
		contentStr = string(rawContent)
	}

	// parse private key
	rawPrivKey, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to sign nostr event")
	}
	privKeyObj, pubKeyObj := btcec.PrivKeyFromBytes(rawPrivKey)
	pubKey := hex.EncodeToString(pubKeyObj.SerializeCompressed()[1:])

	// calculate event id
	obj := []any{
		0,
		pubKey,
		createdAt.Unix(),
		kind,
		tags,
		contentStr,
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to sign nostr event")
	}
	hash := sha256.Sum256(objStr)

	// signature
	sig, err := schnorr.Sign(privKeyObj, hash[:], schnorr.FastSign())
	if err != nil {
		return nil, cerror.Wrap(err, "failed to sign nostr event")
	}

	// return
	return &model.NostrEvent{
		Id:        hex.EncodeToString(hash[:]),
		Pubkey:    pubKey,
		CreatedAt: int(createdAt.Unix()),
		Kind:      kind,
		Tags:      tags,
		Content:   contentStr,
		Sig:       hex.EncodeToString(sig.Serialize()),
	}, nil
}

func NostrVerify(event model.NostrEvent) (bool, error) {
	// check event id
	obj := []any{
		0,
		event.Pubkey,
		event.CreatedAt,
		event.Kind,
		event.Tags,
		event.Content,
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		return false, cerror.Wrap(err, "failed to verify nostr event")
	}
	hash := sha256.Sum256(objStr)
	if hex.EncodeToString(hash[:]) != event.Id {
		return false, cerror.Wrap(cerror.ErrInvalidNostrEventId, "failed to verify nostr event")
	}

	// check signature
	rawPubKey, err := hex.DecodeString(event.Pubkey)
	if err != nil {
		return false, cerror.Wrap(err, "failed to verify nostr event")
	}
	pubKeyObj, err := schnorr.ParsePubKey(rawPubKey)
	if err != nil {
		return false, cerror.Wrap(err, "failed to verify nostr event")
	}
	rawSig, err := hex.DecodeString(event.Sig)
	if err != nil {
		return false, cerror.Wrap(err, "failed to verify nostr event")
	}
	sigObj, err := schnorr.ParseSignature(rawSig)
	if err != nil {
		return false, cerror.Wrap(err, "failed to verify nostr event")
	}
	verified := sigObj.Verify(hash[:], pubKeyObj)
	if !verified {
		return false, cerror.Wrap(cerror.ErrInvalidNostrEventSig, "failed to verify nostr event")
	}
	return true, nil
}
