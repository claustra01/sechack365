package model

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Host           string `json:"host"`
	HashedPassword string `json:"hashed_password"`
	DisplayName    string `json:"display_name"`
	Profile        string `json:"profile"`
	Icon           string `json:"icon"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type ApUserIdentifier struct {
	UserId     string `json:"user_id"`
	Inbox      string `json:"inbox"`
	Outbox     string `json:"outbox"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ApUser struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Host           string `json:"host"`
	HashedPassword string `json:"hashed_password"`
	DisplayName    string `json:"display_name"`
	Profile        string `json:"profile"`
	Icon           string `json:"icon"`
	Inbox          string `json:"inbox"`
	Outbox         string `json:"outbox"`
	PublicKey      string `json:"public_key"`
	PrivateKey     string `json:"private_key"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
