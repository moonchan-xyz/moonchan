package search

import (
	"errors"

	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

type Result struct {
	Accounts []entities.Account `json:"accounts"`
	Statuses []entities.Status  `json:"statues"`
}

func V1() (*Result, error) {
	return nil, errors.New("// TODO")
}

func V2(
	q string,
	t string, // type
	resolve bool,
	following bool,
	accountID string,
	excludeUnreviewed bool,
	maxID, minID string,
	limit int,
	offset int,
) (*Result, error) {
	return nil, errors.New("// TODO")
}
