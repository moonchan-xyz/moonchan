package db

import (
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db = initDB()
)

// init a db
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Init() {
	// Migrate the schema
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Source{})
	db.AutoMigrate(&entities.Role{})

	db.AutoMigrate(&Userdata{})

	db.AutoMigrate(&BlockedBlocking{})
	db.AutoMigrate(&FollowedFollowing{})
	db.AutoMigrate(&entities.Notification{})
}

func init() {
	Init()
}
