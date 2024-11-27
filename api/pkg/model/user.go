package model

type User struct {
	Id             string `json:"id" db:"id"`
	Username       string `json:"username" db:"username"`
	Host           string `json:"host" db:"host"`
	Protocol       string `json:"protocol" db:"protocol"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
	DisplayName    string `json:"display_name" db:"display_name"`
	Profile        string `json:"profile" db:"profile"`
	Icon           string `json:"icon" db:"icon"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
}

type ApUserIdentifier struct {
	UserId     string `json:"user_id" db:"user_id"`
	PublicKey  string `json:"public_key" db:"public_key"`
	PrivateKey string `json:"private_key" db:"private"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}
