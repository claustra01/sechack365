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

func ValidateKeyPair(publicKeyPEM, privateKeyPEM string) error {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "PUBLIC KEY" {
		return cerror.ErrDecodePublicKey
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return cerror.ErrDecodePublicKey
	}

	block, _ = pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return cerror.ErrDecodePrivateKey
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return cerror.ErrDecodePrivateKey
	}

	if publicKey.(*rsa.PublicKey).N.Cmp(privateKey.PublicKey.N) != 0 || publicKey.(*rsa.PublicKey).E != privateKey.PublicKey.E {
		return cerror.ErrInvalidKeyPair
	}

	return nil
}
