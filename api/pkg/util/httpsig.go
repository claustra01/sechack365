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
	switch key := key.(type) {
	case *rsa.PrivateKey:
		b.Type = "RSA PRIVATE KEY"
		//
		b.Bytes = x509.MarshalPKCS1PrivateKey(key)
	case *rsa.PublicKey:
		b.Type = "PUBLIC KEY"
		b.Bytes, err = x509.MarshalPKIXPublicKey(key)
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

func HttpSigSign(keyname string, privKey *rsa.PrivateKey, req *http.Request, content []byte) error {
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

func HttpSigVerify(req *http.Request, content []byte, pubKey *rsa.PublicKey) (string, error) {
	sigHeader := req.Header.Get("Signature")
	if sigHeader == "" {
		return "", cerror.Wrap(cerror.ErrInvalidHttpSig, "failed to verify request")
	}

	// parse headers
	var keyname, algo, heads, b64sig string
	for _, v := range strings.Split(sigHeader, ",") {
		name, val, ok := strings.Cut(v, "=")
		if !ok {
			return "", fmt.Errorf("bad scan: %s from %s", v, sigHeader)
		}
		val = strings.TrimPrefix(val, `"`)
		val = strings.TrimSuffix(val, `"`)
		switch name {
		case "keyId":
			keyname = val
		case "algorithm":
			algo = val
		case "headers":
			heads = val
		case "signature":
			b64sig = val
		default:
			return "", fmt.Errorf("bad sig val: %s from %s", name, sigHeader)
		}
	}
	if keyname == "" || algo == "" || heads == "" || b64sig == "" {
		return "", fmt.Errorf("missing a sig value")
	}

	required := make(map[string]bool)
	required["(request-target)"] = true
	required["date"] = true
	required["host"] = true
	if strings.ToLower(req.Method) != "get" {
		required["digest"] = true
	}
	headers := strings.Split(heads, " ")
	var stuff []string
	for _, h := range headers {
		var s string
		switch h {
		case "(request-target)":
			s = strings.ToLower(req.Method) + " " + req.URL.RequestURI()
		case "host":
			s = req.Host
			if s == "" {
				return "", fmt.Errorf("httpsig: no host header value")
			}
		case "digest":
			s = req.Header.Get(h)
			hash := sha256.New()
			hash.Write(content)
			expv := "SHA-256=" + EncodeBase64(hash.Sum(nil))
			if s != expv {
				return "", fmt.Errorf("digest header '%s' did not match content", s)
			}
		case "date":
			s = req.Header.Get(h)
			d, err := time.Parse(http.TimeFormat, s)
			if err != nil {
				return "", fmt.Errorf("error parsing date header: %s", err)
			}
			now := time.Now()
			if d.Before(now.Add(-30*time.Minute)) || d.After(now.Add(30*time.Minute)) {
				return "", fmt.Errorf("date header '%s' out of range", s)
			}
		default:
			s = req.Header.Get(h)
		}
		delete(required, h)
		stuff = append(stuff, h+": "+s)
	}
	if len(required) > 0 {
		var missing []string
		for h := range required {
			missing = append(missing, h)
		}
		return "", fmt.Errorf("required httpsig headers missing (%s)", strings.Join(missing, ","))
	}

	hash := sha256.New()
	hash.Write([]byte(strings.Join(stuff, "\n")))
	sig := DecodeBase64(b64sig)
	if err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash.Sum(nil), sig); err != nil {
		return keyname, err
	}
	return keyname, nil
}
