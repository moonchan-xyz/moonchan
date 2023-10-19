package entities

import "time"

type Status struct {
	ID                 string            `json:"id"`
	URI                string            `json:"uri"`
	CreatedAt          time.Time         `json:"created_at"`
	Account            *Account          `json:"account"`
	Content            string            `json:"content"`
	Visibility         string            `json:"visibility"`
	Sensitive          bool              `json:"sensitive"`
	SpoilerText        string            `json:"spoiler_text"`
	MediaAttachments   []MediaAttachment `json:"media_attachments"`
	Application        *Application      `json:"application"`
	Mentions           []Mention         `json:"mentions"`
	Tags               []Tag             `json:"tags"`
	Emojis             []interface{}     `json:"emojis"`
	ReblogsCount       int               `json:"reblogs_count"`
	FavouritesCount    int               `json:"favourites_count"`
	RepliesCount       int               `json:"replies_count"`
	URL                string            `json:"url"`
	InReplyToID        interface{}       `json:"in_reply_to_id"`
	InReplyToAccountID interface{}       `json:"in_reply_to_account_id"`
	Reblog             interface{}       `json:"reblog"`
	Poll               *Poll             `json:"poll"`
	Card               *Card             `json:"card"`
	Language           *string           `json:"language"`
	Text               *string           `json:"text"`
	EditedAt           *time.Time        `json:"edited_at"`

	Favourited bool `json:"favourited"`
	Reblogged  bool `json:"reblogged"`
	Muted      bool `json:"muted"`
	Bookmarked bool `json:"bookmarked"`
}
type Mention struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	URL      string `json:"url"`
	Acct     string `json:"acct"`
}
type Tag struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Application struct {
	Name    string  `json:"name"`
	Website *string `json:"website"`
}
type Card struct {
	URL          string     `json:"url"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Type         string     `json:"type"`
	AuthorName   string     `json:"author_name"`
	AuthorURL    string     `json:"author_url"`
	ProviderName string     `json:"provider_name"`
	ProviderURL  string     `json:"provider_url"`
	HTML         string     `json:"html"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Image        *URLstring `json:"image"`
	EmbedURL     URLstring  `json:"embed_url"`
	Blurhash     *string    `json:"blurhash"`
}
