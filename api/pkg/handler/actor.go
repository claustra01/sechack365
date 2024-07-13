package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"

	"github.com/claustra01/sechack365/pkg/framework"
)

type Actor struct {
	Context           []string  `json:"@context"`
	Type              string    `json:"type"`
	Id                string    `json:"id"`
	Inbox             string    `json:"inbox"`
	Outbox            string    `json:"outbox"`
	PreferredUsername string    `json:"preferredUsername"`
	Name              string    `json:"name"`
	Summary           string    `json:"summary"`
	PublicKey         PublicKey `json:"publicKey"`
}

type PublicKey struct {
	Type         string `json:"type"`
	Id           string `json:"id"`
	Owner        string `json:"owner"`
	PublicKeyPem string `json:"publicKeyPem"`
}

func generatePublicKeyPem() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode public key: %w", err)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(publicKeyPEM), nil
}

func MockActor(c *framework.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// mock actor
		url := fmt.Sprintf("https://%s/actor/mock", c.Config.Host)

		publicKeyPem, err := generatePublicKeyPem()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		actor := Actor{
			Context: []string{
				"https://www.w3.org/ns/activitystreams",
				"https://w3id.org/security/v1",
			},
			Type:              "Person",
			Id:                url,
			Inbox:             url + "/inbox",
			Outbox:            url + "/outbox",
			PreferredUsername: "mock",
			Name:              "Mock User",
			Summary:           "The user in activitypub server made by claustra01",
			PublicKey: PublicKey{
				Type:         "Key",
				Id:           url + "#main-key",
				Owner:        url,
				PublicKeyPem: publicKeyPem,
			},
		}

		data, err := json.Marshal(actor)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/activity+json")
		fmt.Fprint(w, string(data))
	}
}
