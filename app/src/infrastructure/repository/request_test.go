package repository_test

import (
	"app/domain"
	"app/infrastructure/repository"
	"log"
	"os"
	"testing"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func setupTestDB(t *testing.T) *gorm.DB {
	log.Print(os.Getenv("DATABASE_URL"))
	// 1. インメモリSQLiteでテスト用DB接続を開く
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	// 2. テーブルをマイグレーション
	if err := db.AutoMigrate(&domain.FriendRequest{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	db.Migrator().DropTable(&domain.FriendRequest{})
	db.AutoMigrate(&domain.FriendRequest{})
	return db
}

// TestCreateAndExists テスト: Createが成功し、その後にExistsがTrueを返すか
func TestCreateAndExists(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.FriendRequestRepositoryImpl{DB: db}

	senderID := "userA"
	targetID := "userB"
	requestID := "uuid-001"
    
	// 1. 作成
	req := domain.SendFriendRequest(senderID, targetID, requestID)
	err := repo.Create(req)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	// 2. 存在確認
	exists, err := repo.Exists(senderID, targetID)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Errorf("Expected Exists to be true, got false")
	}

	// 3. 存在しないリクエストの確認
	exists, _ = repo.Exists("userC", "userD")
	if exists {
		t.Errorf("Expected Exists to be false, got true")
	}
}
