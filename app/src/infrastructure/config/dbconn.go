package config

import (
	"app/infrastructure/models"
	"log"
	"os"
	"testing"

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
		log.Println("failed to connect database", err)
	}

	// グローバル変数に格納
	db = dbconn

	//マイグレーション
	db.AutoMigrate(&models.FriendRequest{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Label{})

	return db
}

// SetupTestDB はテスト用のDB接続をセットアップします
func SetupTestDB(t *testing.T) *gorm.DB {
	// 1. インメモリSQLiteでテスト用DB接続を開く
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect to test database:", err)
	}

	// 2. テーブルをマイグレーション

	db.Migrator().DropTable(&models.FriendRequest{})
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Session{})
	db.Migrator().DropTable(&models.Label{})

	db.AutoMigrate(&models.FriendRequest{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Label{})

	return db
}
