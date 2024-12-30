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
		return "", "", err
	}
	return privKey, pubKey, nil
}

func NostrVerify(pubKey string, digest string, sig string) (bool, error) {
	rawPubKey, err := hex.DecodeString(pubKey)
	if err != nil {
		return false, err
	}
	pubKeyObj, err := schnorr.ParsePubKey(rawPubKey)
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

	// parse private key
	rawPrivKey, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	privKeyObj, _ := btcec.PrivKeyFromBytes(rawPrivKey)

	// calculate digest
	obj := []any{
		0,
		pubKey,
		createdAt.Unix(),
		kind,
		tags,
		string(rawContent),
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(objStr)

	// signature
	sig, err := schnorr.Sign(privKeyObj, hash[:], schnorr.FastSign())
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
