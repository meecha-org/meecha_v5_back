package config

import (
	"app/domain"
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
		log.Fatal("failed to connect database", err)
	}

	// グローバル変数に格納
	db = dbconn

	//マイグレーション
	db.AutoMigrate(&domain.FriendRequest{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Session{})
	db.AutoMigrate(&domain.Label{})

	return db
}

// SetupTestDB はテスト用のDB接続をセットアップします
func SetupTestDB(t *testing.T) *gorm.DB {
	log.Print(os.Getenv("DATABASE_URL"))
	// 1. インメモリSQLiteでテスト用DB接続を開く
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	// 2. テーブルをマイグレーション
	if err := db.AutoMigrate(&domain.FriendRequest{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	db.Migrator().DropTable(&domain.FriendRequest{})
	db.Migrator().DropTable(&domain.User{})
	db.Migrator().DropTable(&domain.Session{})
	db.Migrator().DropTable(&domain.Label{})

	db.AutoMigrate(&domain.FriendRequest{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Session{})
	db.AutoMigrate(&domain.Label{})

	return db
}
