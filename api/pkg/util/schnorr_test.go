package util

import (
	"testing"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
)

func TestNostrSign(t *testing.T) {
	privKey := "3b0038d64fcf4e31cae0c5dc0ab60be7dea5c46c6aa17bf2a526aac3a5cbaa1a"
	createdAt := time.Now()
	kind := 0
	tags := []model.NostrEventTag{}
	content := model.NostrProfile{
		Name:        "test",
		DisplayName: "test",
		About:       "test",
		Picture:     "test.png",
	}
	event, err := NostrSign(privKey, createdAt, kind, tags, content)
	if err != nil {
		t.Errorf("failed to sign nostr event: %v", err)
	}
	if ok, err := NostrVerify(*event); !ok || err != nil {
		t.Errorf("failed to check signed nostr event: %v", err)
	}
}

func TestNostrVerify(t *testing.T) {
	event := model.NostrEvent{
		Id:        "b5ae3e4adade7be6b8993afef0f444284052311ff0dc3f0959b2ef35e1fbafc9",
		Pubkey:    "d5affef57316cde4f40c612ee89ffc54f1e40d3f65e45e4eabc27b5d4542b807",
		CreatedAt: 1735725144,
		Kind:      1,
		Tags:      []model.NostrEventTag{},
		Content:   "test",
		Sig:       "3ce04955da59b3c7c25f1b7044f073a34aa5af490729b5f87e3e29a1323e5d2341641399db3c717d285955f316fe1060fb35d443954b811a2b394d6f6ce969f3",
	}
	if ok, err := NostrVerify(event); !ok || err != nil {
		t.Errorf("failed to verify nostr event: %v", err)
	}
}
