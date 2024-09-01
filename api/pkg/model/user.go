package model

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Host           string `json:"host"`
	Protocol       string `json:"protocol"`
	HashedPassword string `json:"hashed_password"`
	DisplayName    string `json:"display_name"`
	Profile        string `json:"profile"`
	Icon           string `json:"icon"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type ApUserIdentifier struct {
	UserId     string `json:"user_id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
