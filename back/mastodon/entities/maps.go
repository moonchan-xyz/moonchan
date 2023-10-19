package entities

// https://docs.joinmastodon.org/methods/preferences/#get
type Preference struct {
	PostingDefaultVisibility string  `json:"posting:default:visibility"`
	PostingDefaultSensitive  bool    `json:"posting:default:sensitive"`
	PostingDefaultLanguage   *string `json:"posting:default:language"`
	ReadingExpandMedia       string  `json:"reading:expand:media"`
	ReadingExpandSpoilers    bool    `json:"reading:expand:spoilers"`
}
