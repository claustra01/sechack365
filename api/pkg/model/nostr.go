package model

type NostrEvent struct {
	Id        string          `json:"id"`
	Pubkey    string          `json:"pubkey"`
	CreatedAt int             `json:"created_at"` // unix timestamp
	Kind      int             `json:"kind"`
	Tags      []NostrEventTag `json:"tags"`
	Content   string          `json:"content"`
	Sig       string          `json:"sig"`
}

type NostrEventTag []string

type NostrFilter struct {
	// Ids     []string `json:"ids"`
	Authors []string `json:"authors"`
	Kinds   []int    `json:"kinds"`
	// ETags   []string `json:"#e"`
	// PTags   []string `json:"#p"`
	Since int64 `json:"since"` // unix timestamp
	Until int64 `json:"until"` // unix timestamp
	Limit int   `json:"limit"`
}

type NostrProfile struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	About       string  `json:"about"`
	Picture     string  `json:"picture"`
	Nip05       *string `json:"nip05"`
}
