package accounts

import (
	"strconv"

	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

func Get(id string) (account *entities.Account, err error) {
	account = &entities.Account{ID: id}
	account, err = db.ReadAccount(account)
	return
}

func GetCredentials(id string) (account *entities.CredentialAccount, err error) {
	account = &entities.CredentialAccount{Account: entities.Account{ID: id}}
	_, err = db.ReadAccount(&account.Account)
	if err != nil {
		return
	}
	source := &entities.Source{ID: id}
	source, err = db.ReadSource(source)
	if err != nil {
		return
	}
	numID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	// find role
	role := &entities.Role{ID: numID}
	role, err = db.ReadRole(role)
	if err != nil {
		return
	}
	account.Source = source
	account.Role = role
	return
}
