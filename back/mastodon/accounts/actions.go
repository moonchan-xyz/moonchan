package accounts

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func Follow(authorizationID string, id string, reblogs bool, notify bool, languages *string) (*entities.Relationship, error) {
	return nil, nil // todo
}

func Unfollow(authorizationID string, id string) (*entities.Relationship, error) {
	return nil, nil
}
func RemoveFromFollowers(authorizationID string, id string) (*entities.Relationship, error) {
	return nil, nil
}

func Block(authorizationID string, id string) (*entities.Relationship, error) {
	return nil, nil
}

func Unblock(authorizationID string, id string) (*entities.Relationship, error) {
	return nil, nil
}

func Mute(authorizationID string, id string, notifications bool, duration int) (*entities.Relationship, error) {
	return nil, nil
}
func Unmute(authorizationID string, id string) (*entities.Relationship, error) {
	return nil, nil
}
