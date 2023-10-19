package accounts

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func Followers(
	ID,
	maxID,
	sinceID,
	minID string,
	limit int,
) ([]entities.Account, error) {
	return nil, nil //todo
}

func Following(
	ID,
	maxID,
	sinceID,
	minID string,
	limit int,
) ([]entities.Account, error) {
	return nil, nil //todo
}
