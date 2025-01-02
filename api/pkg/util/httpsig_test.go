package util

import (
	"bytes"
	"net/http"
	"testing"
)

func TestConvertRsaPem(t *testing.T) {
	// generate key pair
	privKey, pubKey, err := GenerateRsaKeyPair()
	if err != nil {
		t.Errorf("failed to pem convert test: %v", err)
	}

	// public key enc/dec test
	pubPem, err := EncodePem(pubKey)
	if err != nil {
		t.Errorf("failed to pem convert test: %v", err)
	}
	_, pubDecoded, err := DecodePem(pubPem)
	if err != nil {
		t.Errorf("failed to pem convert test: %v", err)
	}
	if pubDecoded == nil || pubKey.E != pubDecoded.E || pubKey.N.Cmp(pubDecoded.N) != 0 {
		t.Errorf("failed to pem convert test: %v", "not equal public key")
	}

	// private key enc/dec test
	privPem, err := EncodePem(privKey)
	if err != nil {
		t.Errorf("failed to pem convert test: %v", err)
	}
	privDecoded, _, err := DecodePem(privPem)
	if err != nil {
		t.Errorf("failed to pem convert test: %v", err)
	}
	if privDecoded == nil || privKey.E != privDecoded.E || privKey.N.Cmp(privDecoded.N) != 0 {
		t.Errorf("failed to pem convert test: %v", "not equal private key")
	}
}

func TestHttpSigSign(t *testing.T) {
	// generate key pair
	privKey, pubKey, err := GenerateRsaKeyPair()
	if err != nil {
		t.Errorf("failed to sign request test: %v", err)
	}

	// sign test
	body := []byte("test")
	req, err := http.NewRequest("POST", "https://example.com", bytes.NewReader(body))
	if err != nil {
		t.Errorf("failed to sign request test: %v", err)
	}
	keyId := "https://localhost/api/v1/users/test#main-key"
	if err := HttpSigSign(keyId, privKey, req, body); err != nil {
		t.Errorf("failed to sign request test: %v", err)
	}

	// verify test
	keyname, err := HttpSigVerify(req, body, pubKey)
	if err != nil {
		t.Errorf("failed to verify request test: %v", err)
	}
	if keyname != keyId {
		t.Errorf("failed to verify request test: %v", "not equal key id")
	}
}
