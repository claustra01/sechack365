package model

type User struct {
	Id                string `json:"id"`
	UserId            string `json:"user_id"`
	Host              string `json:"host"`
	EncryptedPassword string `json:"encrypted_password"`
	DisplayName       string `json:"display_name"`
	Profile           string `json:"profile"`
}

type ApUserIdentifier struct {
	UserId     string `json:"user_id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type ApUser struct {
	Id                string `json:"id"`
	UserId            string `json:"user_id"`
	Host              string `json:"host"`
	EncryptedPassword string `json:"encrypted_password"`
	DisplayName       string `json:"display_name"`
	Profile           string `json:"profile"`
	PublicKey         string `json:"public_key"`
	PrivateKey        string `json:"private_key"`
}
