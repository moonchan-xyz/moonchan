package methods

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func GetFollowRequests(id string, maxID, sinceID string, limit int) ([]entities.Account, error) {
	return nil, nil // todo
}

func AcceptFollowRequest(id, accountID string) (*entities.Relationship, error) {
	return nil, nil // todo

}

func RejectFollowRequest(id, accountID string) (*entities.Relationship, error) {
	return nil, nil // todo
}
