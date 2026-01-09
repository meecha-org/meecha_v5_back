package config

import (
	"app/domain"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	db *gorm.DB = nil
)

func Init() *gorm.DB {
	// MySQL接続
	dbconn, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// グローバル変数に格納
	db = dbconn

	//マイグレーション
	db.AutoMigrate(&domain.FriendRequest{})

	return db
}
