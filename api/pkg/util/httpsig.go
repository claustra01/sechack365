package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"
	"time"

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

func SignRequest(keyname string, privKey *rsa.PrivateKey, req *http.Request, content []byte) error {
	// headers
	headers := []string{"(request-target)", "date", "host"}
	if strings.ToLower(req.Method) != "get" {
		headers = append(headers, "content-type", "digest")
	}

	// digest
	var stuff []string
	for _, h := range headers {
		var s string
		switch h {
		case "(request-target)":
			s = strings.ToLower(req.Method) + " " + req.URL.RequestURI()
		case "date":
			s = req.Header.Get(h)
			if s == "" {
				s = time.Now().UTC().Format(http.TimeFormat)
				req.Header.Set(h, s)
			}
		case "host":
			s = req.Header.Get(h)
			if s == "" {
				s = req.URL.Hostname()
				req.Header.Set(h, s)
			}
		case "content-type":
			s = req.Header.Get(h)
		case "digest":
			s = req.Header.Get(h)
			if s == "" {
				hash := sha256.New()
				hash.Write(content)
				s = "SHA-256=" + EncodeBase64(hash.Sum(nil))
				req.Header.Set(h, s)
			}
		}
		stuff = append(stuff, h+": "+s)
	}

	// signature
	hash := sha256.New()
	hash.Write([]byte(strings.Join(stuff, "\n")))
	sig, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return cerror.Wrap(err, "failed to sign request")
	}
	b64sig := EncodeBase64(sig)
	sigHeader := fmt.Sprintf(`keyId="%s",algorithm="%s",headers="%s",signature="%s"`, keyname, "rsa-sha256", strings.Join(headers, " "), b64sig)
	req.Header.Set("Signature", sigHeader)
	return nil
}
