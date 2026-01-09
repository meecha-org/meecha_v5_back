package repository_test

import (
	"app/domain"
	"app/infrastructure/config"
	"app/infrastructure/repository"
	"testing"

)

// TestCreateAndExists テスト: Createが成功し、その後にExistsがTrueを返すか
func TestCreateAndExists(t *testing.T) {

	db := config.SetupTestDB(t)
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
