package util

import (
	"testing"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

func TestVerify(t *testing.T) {
	id := "e60d885074b910cc93d63be506d95144661ef83b4e45cc00cf6d217ca41eb95e"
	pubKey := "f81611363554b64306467234d7396ec88455707633f54738f6c4683535098cd3"
	sig := "6b83ad045de1aac986566e34f63efe53104f9892c001869c3e720515ad6cc589dfd22cf4a04e93223135de815e545d4f8446848aa323fcaf1060e32b72dd8aa8"
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
	event, err := NostrSign(privKey, createdAt, kind, tags, content)
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
