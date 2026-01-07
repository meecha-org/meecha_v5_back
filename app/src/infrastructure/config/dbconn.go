package config

import (
	"app/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	db *gorm.DB = nil
)

func Init() *gorm.DB {
	// PostgreSQL接続
	dbconn, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// グローバル変数に格納
	db = dbconn

	//マイグレーション
	db.AutoMigrate(&domain.FriendRequest{})

	return db
}
