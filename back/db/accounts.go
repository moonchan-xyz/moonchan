package db

import (
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

func CreateAccount(account *entities.Account) error {
	tx := db.Create(account)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadAccount(account *entities.Account) (*entities.Account, error) {
	tx := db.Take(account)
	if tx.Error != nil {
		return account, tx.Error
	}
	return account, nil
}

func ReadAccountByAcct(acct string) (*entities.Account, error) {
	account := &entities.Account{Acct: acct}
	return ReadAccount(account)
}

func UpdateAccount(account *entities.Account) (*entities.Account, error) {
	tx := db.Save(account)
	if tx.Error != nil {
		return account, tx.Error
	}
	return account, nil
}

func CreateSource(source *entities.Source) error {
	tx := db.Create(source)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadSource(source *entities.Source) (*entities.Source, error) {
	tx := db.Take(source)
	if tx.Error != nil {
		return source, tx.Error
	}
	return source, nil
}

func CreateRole(role *entities.Role) error {
	tx := db.Create(role)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadRole(role *entities.Role) (*entities.Role, error) {
	tx := db.Take(role)
	if tx.Error != nil {
		return role, tx.Error
	}
	return role, nil
}

type BlockedBlocking struct {
	BlockedID  string `gorm:"primaryKey"`
	BlockingID string `gorm:"primaryKey"`
}

func CreateRelationBlock(blockingID, BlockedID string) error {
	relation := &BlockedBlocking{BlockingID: blockingID, BlockedID: BlockedID}
	tx := db.Create(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadRelationBlock(blockingID, BlockedID string) error {
	relation := &BlockedBlocking{BlockingID: blockingID, BlockedID: BlockedID}
	tx := db.Take(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteRelationBlock(blockingID, BlockedID string) error {
	relation := &BlockedBlocking{BlockingID: blockingID, BlockedID: BlockedID}
	tx := db.Delete(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

type FollowedFollowing struct {
	FollowedID  string `gorm:"primaryKey"`
	FollowingID string `gorm:"primaryKey"`
}

func CreateRelationFollow(followingID, followedID string) error {
	relation := &FollowedFollowing{FollowingID: followingID, FollowedID: followedID}
	tx := db.Create(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadRelationFollow(followingID, followedID string) error {
	relation := &FollowedFollowing{FollowingID: followingID, FollowedID: followedID}
	tx := db.Take(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteRelationFollow(followingID, followedID string) error {
	relation := &FollowedFollowing{FollowingID: followingID, FollowedID: followedID}
	tx := db.Delete(relation)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// list
func ReadRelationFollowedBy(followingID string, limit, offset int) (accounts []entities.Account, err error) {
	sql := `
	SELECT accounts.* FROM 
	relation_follows INNER JOIN accounts
	ON relation_follows.followed_id = accounts.id
	WHERE relation_follows.following_id = ?
	LIMIT ? OFFSET ?
	`
	tx := db.Raw(sql, followingID, limit, offset).Scan(&accounts)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}

func ReadRelationFollowingBy(followedID string, limit, offset int) (accounts []entities.Account, err error) {
	sql := `
	SELECT accounts.* FROM 
	relation_follows INNER JOIN accounts
	ON relation_follows.following_id = accounts.id
	WHERE relation_follows.followed_id = ?
	LIMIT ? OFFSET ?
	`
	tx := db.Raw(sql, followedID, limit, offset).Scan(&accounts)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
