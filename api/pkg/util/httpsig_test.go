package util

import (
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
