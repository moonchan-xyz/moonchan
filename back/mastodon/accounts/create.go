package accounts

import (
	"os"
	"strconv"

	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
	"github.com/moonchan-xyz/moonchan/back/utils"
)

// https://docs.joinmastodon.org/methods/accounts/#create
func Create(
	username,
	email,
	password string,
	aggrement bool,
	locale,
	reason string,
) error {
	userdata := &db.Userdata{
		Email:  email,
		Passwd: utils.Hash(password),
		Reason: reason,
	}
	err := db.CreateUserdata(userdata)
	if err != nil {
		return err
	}

	host := os.Getenv("HOST")
	id := utils.Timestamp()
	var createdAt entities.DATEstring
	createdAt.FromTimestamp(id)

	account := &entities.Account{
		ID:        strconv.Itoa(int(id)),
		Username:  username,
		Acct:      utils.ParseAcct(username, host),
		CreatedAt: createdAt,
	}
	err = db.CreateAccount(account)
	if err != nil {
		return err
	}
	return err
}
