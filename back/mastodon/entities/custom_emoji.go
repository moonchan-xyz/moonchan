// https://docs.joinmastodon.org/entities/CustomEmoji/

package entities

type CustomEmoji struct {
	Shortcode       string    `json:"shortcode"`
	Url             URLstring `json:"url"`
	StaticUrl       URLstring `json:"static_url"`
	VisibleInPicker bool      `json:"visible_in_picker"`
	Category        string    `json:"category"`
}
