package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/claustra01/sechack365/pkg/cerror"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, cerror.Wrap(err, "failed to generate rsa key pair")
	}
	return privKey, &privKey.PublicKey, nil
}

func EncodePem(key any) (string, error) {
	var b pem.Block
	var err error
	switch key.(type) {
	case *rsa.PrivateKey:
		b.Type = "RSA PRIVATE KEY"
		b.Bytes = x509.MarshalPKCS1PrivateKey(key.(*rsa.PrivateKey))
	case *rsa.PublicKey:
		b.Type = "PUBLIC KEY"
		b.Bytes, err = x509.MarshalPKIXPublicKey(key.(*rsa.PublicKey))
		if err != nil {
			return "", cerror.Wrap(err, "failed to encode pem")
		}
	default:
		return "", cerror.Wrap(cerror.ErrUnknownKeyType, "failed to encode pem")
	}
	return string(pem.EncodeToMemory(&b)), nil
}

func DecodePem(p string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	b, _ := pem.Decode([]byte(p))
	if b == nil {
		return nil, nil, cerror.Wrap(cerror.ErrInvalidKeyPem, "failed to decode pem")
	}
	switch b.Type {
	case "RSA PRIVATE KEY":
		privKey, err := x509.ParsePKCS1PrivateKey(b.Bytes)
		if err != nil {
			return nil, nil, cerror.Wrap(err, "failed to decode pem")
		}
		return privKey, &privKey.PublicKey, nil
	case "PUBLIC KEY":
		pubKey, err := x509.ParsePKIXPublicKey(b.Bytes)
		if err != nil {
			return nil, nil, cerror.Wrap(err, "failed to decode pem")
		}
		return nil, pubKey.(*rsa.PublicKey), nil
	default:
		return nil, nil, cerror.Wrap(cerror.ErrUnknownKeyType, "failed to decode pem")
	}
}

func GenerateKeyPemPair() (string, string, error) {
	privKey, pubKey, err := GenerateRsaKeyPair()
	if err != nil {
		return "", "", cerror.Wrap(err, "failed to generate key pem pair")
	}
	privKeyPem, err := EncodePem(privKey)
	if err != nil {
		return "", "", cerror.Wrap(err, "failed to generate key pem pair")
	}
	pubKeyPem, err := EncodePem(pubKey)
	if err != nil {
		return "", "", cerror.Wrap(err, "failed to generate key pem pair")
	}
	return privKeyPem, pubKeyPem, nil
}
