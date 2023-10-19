package profile

import (
	"os"

	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

func DeleteAvatar(id string) (*entities.CredentialAccount, error) {
	var DEFAULT_AVATAR = os.Getenv("DEFAULT_AVATAR")
	account := &entities.CredentialAccount{Account: entities.Account{ID: id}}
	_, err := db.ReadAccount(&account.Account)
	if err != nil {
		return account, err
	}
	// delete avatar
	account.Avatar = entities.URLstring(DEFAULT_AVATAR)
	account.AvatarStatic = entities.URLstring(DEFAULT_AVATAR)
	_, err = db.UpdateAccount(&account.Account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func DeleteHeader(id string) (*entities.CredentialAccount, error) {
	var DEFAULT_HEADER = os.Getenv("DEFAULT_HEADER")
	account := &entities.CredentialAccount{Account: entities.Account{ID: id}}
	_, err := db.ReadAccount(&account.Account)
	if err != nil {
		return account, err
	}
	// delete header
	account.Header = entities.URLstring(DEFAULT_HEADER)
	account.HeaderStatic = entities.URLstring(DEFAULT_HEADER)
	_, err = db.UpdateAccount(&account.Account)
	if err != nil {
		return account, err
	}

	return account, nil
}
