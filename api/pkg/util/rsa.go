package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/claustra01/sechack365/pkg/cerror"
)

func GenerateKeyPair() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", cerror.ErrGenerateRsaKey
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", cerror.ErrEncodePublicKey
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	if privateKeyBytes == nil {
		return "", "", cerror.ErrEncodePrivateKey
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return string(publicKeyPEM), string(privateKeyPEM), nil
}

func DecodePublicKeyPem(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, cerror.ErrDecodePublicKey
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, cerror.ErrDecodePublicKey
	}

	return publicKey.(*rsa.PublicKey), nil
}

func DecodePrivateKeyPem(privateKeyPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, cerror.ErrDecodePrivateKey
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, cerror.ErrDecodePrivateKey
	}

	return privateKey, nil
}

func ValidateKeyPair(publicKeyPEM, privateKeyPEM string) error {
	publicKey, err := DecodePublicKeyPem(publicKeyPEM)
	if err != nil {
		return err
	}

	privateKey, err := DecodePrivateKeyPem(privateKeyPEM)
	if err != nil {
		return err
	}

	if publicKey.N.Cmp(privateKey.PublicKey.N) != 0 || publicKey.E != privateKey.PublicKey.E {
		return cerror.ErrInvalidKeyPair
	}

	return nil
}
