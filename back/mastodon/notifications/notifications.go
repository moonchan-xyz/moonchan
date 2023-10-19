package notifications

import (
	"errors"

	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

// https://docs.joinmastodon.org/methods/notifications/#get
// user's id
func Get(
	id string,
	maxID, sinceID, minID string,
	limit int,
	types []string,
	excludeTypes []string,
	accountID string,
) ([]entities.Notification, error) {

	return nil, errors.New("// TODO")
}

// notification's id
func GetOne(id string) (*entities.Notification, error) {
	notification := &entities.Notification{ID: id}
	notification, err := db.ReadNotification(notification)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

// user's id
func Clear(id string) error {
	err := db.SetNotificationsRead([]string{id})
	if err != nil {
		return err
	}
	return nil
}
