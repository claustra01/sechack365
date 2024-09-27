package service

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/claustra01/sechack365/pkg/model"
	"github.com/claustra01/sechack365/pkg/util"
)

type NostrService struct {
	Ws model.IWsHandler
}

func (s *NostrService) Req(id string, filter model.NostrFilter) ([]string, error) {
	var arr []any
	arr = append(arr, "REQ")
	arr = append(arr, id)
	arr = append(arr, filter)
	reqMsg, err := json.Marshal(arr)
	if err != nil {
		return nil, err
	}
	if err := s.Ws.Send(string(reqMsg)); err != nil {
		return nil, err
	}

	var msgs []string
	for {
		resMsg, err := s.Ws.Receive()
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(resMsg, `["EOSE",`) {
			break
		}
		msgs = append(msgs, resMsg)
	}
	return msgs, nil
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
	msgs, err := s.Req(id.String(), filter)
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
