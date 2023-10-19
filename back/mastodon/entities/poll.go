package entities

import "time"

type Poll struct {
	ID          string        `json:"id"`
	ExpiresAt   *time.Time    `json:"expires_at"`
	Expired     bool          `json:"expired"`
	Multiple    bool          `json:"multiple"`
	VotesCount  int           `json:"votes_count"`
	VotersCount *int          `json:"voters_count"`
	Options     []Option      `json:"options"`
	Emojis      []CustomEmoji `json:"emojis"`
	Voted       *bool         `json:"voted"`
	OwnVotes    []int         `json:"own_votes"`
}
type Option struct {
	Title      string `json:"title"`
	VotesCount *int   `json:"votes_count"`
}
