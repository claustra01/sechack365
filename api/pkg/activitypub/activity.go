package activitypub

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/claustra01/sechack365/pkg/cerror"
)

type SignParams struct {
	Host       string
	KeyId      string
	PrivateKey *rsa.PrivateKey
}

type FollowActivity struct {
	Context any    `json:"@context"`
	Type    string `json:"type"`
	Id      string `json:"id"`
	Actor   string `json:"actor"`
	Object  string `json:"object"`
}

func SendActivity(url string, activity any, sigParams SignParams) ([]byte, error) {
	reqBody, err := json.Marshal(activity)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	signedDate := time.Now().Format(http.TimeFormat)

	req.Header.Set("Host", sigParams.Host)
	req.Header.Set("Date", signedDate)
	req.Header.Set("Content-Type", "application/activity+json")

	hash := sha256.Sum256(reqBody)
	digest := base64.StdEncoding.EncodeToString(hash[:])
	digestHeader := fmt.Sprintf("SHA-256=%s", digest)
	req.Header.Set("Digest", digestHeader)

	// TODO: このsigningStringを署名するのが正しいという認識だが理解が怪しいので後日実装する
	// signingString := fmt.Sprintf("(request-target): post %s\nhost: %s\ndate: %s\ndigest: %s", url, sigParams.Host, signedDate, digestHeader)
	rawSign, err := rsa.SignPKCS1v15(rand.Reader, sigParams.PrivateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}
	encodedSign := base64.StdEncoding.EncodeToString(rawSign)
	signatureHeader := fmt.Sprintf(`keyId="%s",algorithm="rsa-sha256",headers="(request-target) host date digest",signature="%s"`, sigParams.KeyId, encodedSign)
	req.Header.Set("Signature", signatureHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror.Wrap(cerror.ErrPushActivity, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, cerror.Wrap(cerror.ErrPushActivity, string(body))
	}
	return body, nil
}

func BuildFollowActivitySchema(id, host, followerId, followeeUrl string) *FollowActivity {
	// object is followee actor
	return &FollowActivity{
		// Context: ApContext[:],
		Type:   "Follow",
		Id:     fmt.Sprintf("https://%s/follows/%s", host, id),
		Actor:  BuildActorUrl(host, followerId),
		Object: followeeUrl,
	}
}
