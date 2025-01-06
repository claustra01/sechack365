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
	Authors []string `json:"authors,omitempty"`
	Kinds   []int    `json:"kinds,omitempty"`
	PTags   []string `json:"#p,omitempty"`
	Since   int64    `json:"since,omitempty"` // unix timestamp
	Until   int64    `json:"until,omitempty"` // unix timestamp
	Limit   int      `json:"limit,omitempty"`
}

type NostrProfile struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	About       string `json:"about"`
	Picture     string `json:"picture"`
}
