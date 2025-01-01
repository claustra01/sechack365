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

func (s *NostrService) sendReq(id string, filter model.NostrFilter) ([]string, error) {
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

func (s *NostrService) sendEvent(event model.NostrEvent) error {
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

func (s *NostrService) GetRemoteProfile(pubkey string) (*model.NostrProfile, error) {
	id := util.NewUuid()
	filter := model.NostrFilter{
		Kinds:   []int{0},
		Authors: []string{pubkey},
		Limit:   1,
		Since:   0,
		Until:   time.Now().Unix(),
	}
	msgs, err := s.sendReq(id.String(), filter)
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

func (s *NostrService) GetRemotePosts(pubKeys []string, since time.Time) ([]*model.NostrEvent, error) {
	id := util.NewUuid()
	filter := model.NostrFilter{
		Kinds:   []int{1},
		Authors: pubKeys,
		// TODO: Limit should be configurable
		Limit: 100,
		Since: since.Unix() + 1,
		Until: time.Now().Unix(),
	}
	msgs, err := s.sendReq(id.String(), filter)
	if err != nil {
		return []*model.NostrEvent{}, err
	}
	if len(msgs) == 0 {
		return []*model.NostrEvent{}, nil
	}
	var events []*model.NostrEvent
	for _, msg := range msgs {
		var arr []any
		if err := json.Unmarshal([]byte(msg), &arr); err != nil {
			return nil, err
		}
		eventRaw, ok := arr[2].(map[string]any)
		if !ok {
			continue
		}
		event := model.NostrEvent{
			Id:        eventRaw["id"].(string),
			Pubkey:    eventRaw["pubkey"].(string),
			CreatedAt: int(eventRaw["created_at"].(float64)),
			Kind:      int(eventRaw["kind"].(float64)),
			// TODO: Tags should be parsed
			Tags:    []model.NostrEventTag{},
			Content: eventRaw["content"].(string),
			Sig:     eventRaw["sig"].(string),
		}
		events = append(events, &event)
	}
	return events, nil
}

func (s *NostrService) PublishProfile(privKey string, profile *model.NostrProfile) error {
	event, err := util.NostrSign(privKey, time.Now(), 0, []model.NostrEventTag{}, profile)
	if err != nil {
		return err
	}
	if err := s.sendEvent(*event); err != nil {
		return err
	}
	return nil
}

func (s *NostrService) PublishPost(privKey string, note string) error {
	event, err := util.NostrSign(privKey, time.Now(), 1, []model.NostrEventTag{}, note)
	if err != nil {
		return err
	}
	if err := s.sendEvent(*event); err != nil {
		return err
	}
	return nil
}

func (s *NostrService) PublishFollow(privKey string, pubKeys []string) error {
	followList := []model.NostrEventTag{}
	for _, pubKey := range pubKeys {
		// NOTE: "p", pubKey, mainRelay, petName (Ref: NIP-02)
		followList = append(followList, model.NostrEventTag{"p", pubKey, "", ""})
	}
	event, err := util.NostrSign(privKey, time.Now(), 3, followList, "")
	if err != nil {
		return err
	}
	if err := s.sendEvent(*event); err != nil {
		return err
	}
	return nil
}
