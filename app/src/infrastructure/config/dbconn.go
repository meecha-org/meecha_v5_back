package config

import (
	"app/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	db *gorm.DB = nil
)
func Init() *gorm.DB {

	//※コンテナ確率次第直します
	dsn := "host=localhost user=meecha password=meecha_pass name=meecha port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	// PostgreSQL接続
	dbconn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// グローバル変数に格納
	db = dbconn

	//マイグレーション
	db.AutoMigrate(&domain.FriendRequest{})

	return db
}
