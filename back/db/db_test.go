package db

import (
	"testing"

	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

func TestCreateAccount(t *testing.T) {
	Init()
	tx := db.Create(&entities.Account{
		ID:             "",
		Username:       "",
		Acct:           "",
		Url:            "",
		DisplayName:    "",
		Note:           "",
		Avatar:         "",
		AvatarStatic:   "",
		Header:         "",
		HeaderStatic:   "",
		Locked:         false,
		Fields:         nil,
		Emojis:         nil,
		Bot:            false,
		Group:          false,
		Discoverable:   nil,
		Noindex:        nil,
		Moved:          nil,
		Suspended:      false,
		Limited:        false,
		CreatedAt:      "",
		LastStatusAt:   nil,
		StatusesCount:  0,
		FollowersCount: 0,
		FollowingCount: 0,
	})
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}
