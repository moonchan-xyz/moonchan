// https://docs.joinmastodon.org/entities/Account/

// info:
// *(string|XXstring) means nullable

// log:
// username 不能加 gorm:"index:username,unique,type:text collate nocase" 吧，只有username的
// 后面那个用";"还是用" "还是用","分开的啊
// 后面那个叫什么

package entities

type Account struct {
	ID             string        `json:"id" gorm:"primaryKey"`
	Username       string        `json:"username"`
	Acct           string        `json:"acct" gorm:"unique_index type:text collate nocase"` // why "index:acct,unique,type:text collate nocase" not work
	Url            URLstring     `json:"url"`
	DisplayName    string        `json:"display_name"`
	Note           HTMLstring    `json:"note"`
	Avatar         URLstring     `json:"avatar"`
	AvatarStatic   URLstring     `json:"avatar_static"`
	Header         URLstring     `json:"header"`
	HeaderStatic   URLstring     `json:"header_static"`
	Locked         bool          `json:"locked"`
	Fields         []Field       `json:"fields" gorm:"type:bytes;serializer:json"` //
	Emojis         []CustomEmoji `json:"emojis" gorm:"type:bytes;serializer:json"` //
	Bot            bool          `json:"bot"`
	Group          bool          `json:"group"`
	Discoverable   *bool         `json:"discoverable"`
	Noindex        *bool         `json:"noindex"`
	Moved          *Account      `json:"moved" gorm:"type:bytes;serializer:json"`
	Suspended      bool          `json:"suspended"` // false if not set
	Limited        bool          `json:"limited"`   // false if not set
	CreatedAt      DATEstring    `json:"created_at"`
	LastStatusAt   *DATEstring   `json:"last_status_at"`
	StatusesCount  int           `json:"statuses_count"`
	FollowersCount int           `json:"followers_count"`
	FollowingCount int           `json:"following_count"`
	// // CredentialAccount entity attributes
	// Source *Source `json:"source,omitempty" gorm:"-"`
	// Role   *Role   `json:"role,omitempty" gorm:"-"`
	// // MutedAccount entity attributes
	// MuteExpiresAt *string `json:"mute_expires_at"`
}

type CredentialAccount struct {
	Account
	Source *Source `json:"source,omitempty" gorm:"-"`
	Role   *Role   `json:"role,omitempty" gorm:"-"`
}

// https://docs.joinmastodon.org/entities/Account/#Field
type Field struct {
	// Field entity attributes
	Name       string      `json:"name"`
	Value      HTMLstring  `json:"value"`
	VerifiedAt *DATEstring `json:"verified_at"`
}

// https://docs.joinmastodon.org/entities/Account/#CredentialAccount
type Source struct {
	ID                  string  `json:"-"`
	Privacy             string  `json:"privacy"`
	Sensitive           bool    `json:"sensitive"`
	Language            string  `json:"language"`
	Note                string  `json:"note"`
	Fields              []Field `json:"fields" gorm:"type:bytes;serializer:json"`
	FollowRequestsCount int     `json:"follow_requests_count"`
}

// https://docs.joinmastodon.org/entities/Role/
type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Permissions int    `json:"permissions"`
	Highlighted bool   `json:"highlighted"`
}
