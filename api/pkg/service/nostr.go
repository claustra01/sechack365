package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type NostrService struct {
	Ws model.IWsHandler
}

func (s *NostrService) req(id string, filter model.NostrFilter) ([]string, error) {
	reqObj := []any{"REQ", id, filter}
	reqMsg, err := json.Marshal(reqObj)
	if err != nil {
		return nil, cerror.Wrap(err, "failed to request nostr event")
	}
	if err := s.Ws.Send(string(reqMsg)); err != nil {
		return nil, cerror.Wrap(err, "failed to request nostr event")
	}

	var msgs []string
	for {
		resMsg, err := s.Ws.Receive()
		if err != nil {
			return nil, cerror.Wrap(err, "failed to request nostr event")
		}
		var resObj []any
		if err := json.Unmarshal([]byte(resMsg), &resObj); err != nil {
			return nil, cerror.Wrap(err, "failed to request nostr event")
		}
		if resObj[0] == "EOSE" {
			break
		}
		msgs = append(msgs, resMsg)
	}
	return msgs, nil
}

func (s *NostrService) event(event model.NostrEvent) error {
	eventObj := []any{"EVENT", event}
	eventMsg, err := json.Marshal(eventObj)
	if err != nil {
		return cerror.Wrap(err, "failed to post nostr event")
	}
	if err := s.Ws.Send(string(eventMsg)); err != nil {
		return cerror.Wrap(err, "failed to post nostr event")
	}

	resMsg, err := s.Ws.Receive()
	if err != nil {
		return cerror.Wrap(err, "failed to post nostr event")
	}
	var resObj []any
	if err := json.Unmarshal([]byte(resMsg), &resObj); err != nil {
		return cerror.Wrap(err, "failed to post nostr event")
	}
	if resObj[0] != "OK" {
		return cerror.Wrap(cerror.ErrNostrRelayResNotOk, "failed to post nostr event")
	}
	if resObj[2] != true {
		return cerror.Wrap(fmt.Errorf("%v", resObj[3]), "failed to post nostr event")
	}
	return nil
}

func (s *NostrService) GetUserProfile(pubkey string) (*model.NostrProfile, error) {
	id := util.NewUuid()
	filter := model.NostrFilter{
		Kinds:   []int{0},
		Authors: []string{pubkey},
		Limit:   1,
		Since:   0,
		Until:   time.Now().Unix(),
	}
	msgs, err := s.req(id.String(), filter)
	if err != nil {
		return nil, err
	}
	if len(msgs) == 0 {
		return nil, nil
	}
	var arr []any
	if err := json.Unmarshal([]byte(msgs[0]), &arr); err != nil {
		return nil, err
	}
	obj, ok := arr[2].(map[string]any)
	if !ok {
		return nil, nil
	}
	var profile model.NostrProfile
	content, ok := obj["content"].(string)
	if !ok {
		return nil, nil
	}
	content = strings.Replace(content, `\"`, `"`, -1)
	if err := json.Unmarshal([]byte(content), &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (s *NostrService) PostUserProfile(privKey string, pubKey string, profile *model.NostrProfile) error {
	event, err := util.NostrSign(privKey, pubKey, time.Now(), 0, []model.NostrEventTag{}, profile)
	if err != nil {
		return err
	}
	if err := s.event(*event); err != nil {
		return err
	}
	return nil
}
