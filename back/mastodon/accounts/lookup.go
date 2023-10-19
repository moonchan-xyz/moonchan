// TODO

package accounts

import (
	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

// https://docs.joinmastodon.org/methods/accounts/#lookup
func LookUp(acct string) (*entities.Account, error) {
	account, err := db.ReadAccountByAcct(acct)
	if err != nil {
		return account, err
	}
	// TODO 去查询远程
	return account, err
}
