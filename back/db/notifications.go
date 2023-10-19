package db

import "github.com/moonchan-xyz/moonchan/back/mastodon/entities"

func CreateNotification(notification *entities.Notification) error {
	tx := db.Create(notification)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadNotification(notification *entities.Notification) (*entities.Notification, error) {
	tx := db.Take(notification)
	if tx.Error != nil {
		return notification, tx.Error
	}
	return notification, nil
}

func UpdateNotification(notification *entities.Notification) (*entities.Notification, error) {
	tx := db.Save(notification)
	if tx.Error != nil {
		return notification, tx.Error
	}
	return notification, nil
}

func SetNotificationsRead(id []string) error {
	tx := db.Model(&entities.Notification{}).Where("(id) IN ?", id).Update("read", true)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
