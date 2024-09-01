package activitypub

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/claustra01/sechack365/pkg/cerror"
)

type FollowActivity struct {
	Context any    `json:"@context"`
	Type    string `json:"type"`
	Id      string `json:"id"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

// TODO: http signature
func SendActivity(url string, activity any, keyId string, privateKey *rsa.PrivateKey) ([]byte, error) {
	reqBody, err := json.Marshal(activity)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/activity+json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, cerror.ErrPushActivity
	}
	return body, nil
}

func BuildFollowActivitySchema(id, followerName, followerHost, followeeUrl string) *FollowActivity {
	// object is followee actor
	return &FollowActivity{
		Context: ApContext[:],
		Type:    "Follow",
		Id:      fmt.Sprintf("https://%s/follows/%s", followerHost, id),
		Actor:   BuildActorUrl(followerHost, followerName),
		Object:  followeeUrl,
	}
}
