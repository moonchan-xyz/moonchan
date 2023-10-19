package entities

type MediaAttachment struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	URL         string         `json:"url"`
	PreviewURL  string         `json:"preview_url"`
	RemoteURL   *string        `json:"remote_url"`
	Meta        map[string]any `json:"meta"`
	Description string         `json:"description"`
	Blurhash    string         `json:"blurhash"`
	// TextURL     string         `json:"text_url"` // Deprecated
}
