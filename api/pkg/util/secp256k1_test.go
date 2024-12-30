package util

import (
	"testing"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

func TestVerify(t *testing.T) {
	id := "2eb85ddf6ceb5dece487a250ed98993fba5e94f69a88d3438e2afb42a8c9d729"
	pubKey := "48b8d2a8f1fc8784e02919a194cc9d2f498a56b5a2aa59476ddf14d76ee536cc"
	sig := "63d45c3e09d8026e357951bd86bf1dfe5acfa61f704f5bc032bb5622a59a7539ef234383bbbda2b06275e51fc466d1a690890cb8c8a1923fbd8e472d12002254"
	ok, err := NostrVerify(pubKey, id, sig)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("failed to verify")
	}
}

// FIXME: random fail
func TestSign(t *testing.T) {
	// sign
	privKey, pubKey, err := GenerateNostrKeyPair()
	if err != nil {
		t.Error(err)
	}
	createdAt := time.Now()
	kind := 0
	tags := []model.NostrEventTag{}
	content := model.NostrProfile{
		Name: "test",
	}
	event, err := NostrSign(privKey, pubKey, createdAt, kind, tags, content)
	if err != nil {
		t.Error(err)
	}
	if event == nil {
		t.Error("failed to sign")
	}

	// verify
	ok, err := NostrVerify(pubKey, event.Id, event.Sig)
	if err != nil {
		t.Error(err)
	}
	if !ok {
		t.Error("failed to verify")
	}
}
