package domain_test

import (
	"testing"
	"time"
    
	"app/domain"
)

// TestUserCreation テスト: Userエンティティが正しく作成されるか
func TestUserCreation(t *testing.T) {
	UserID := "user123"
	Name := "Test User"
	Email := "test@example.com"
	ProvCode := domain.Google
	ProvUID := "google_uid_123"
	PasswordHash := "hashed_password"
	CreatedAt := time.Now().Unix()
	IsBanned := 0
	IsSystem := 0
	UpdatedAt := time.Now().Unix()

	req := domain.NewUser(UserID, Name, Email, ProvUID, PasswordHash, ProvCode)

	// 1. 各フィールドが正しく設定されているか確認
	if req.UserID != UserID {
		t.Errorf("Expected UserID %s, got %s", UserID, req.UserID)
	}
	if req.Name != Name {
		t.Errorf("Expected Name %s, got %s", Name, req.Name)
	}
	if req.Email != Email {
		t.Errorf("Expected Email %s, got %s", Email, req.Email)
	}
	if req.ProvCode != ProvCode {
		t.Errorf("Expected ProvCode %s, got %s", ProvCode, req.ProvCode)
	}
	if req.ProvUID != ProvUID {
		t.Errorf("Expected ProvUID %s, got %s", ProvUID, req.ProvUID)
	}
	if req.PasswordHash != PasswordHash {
		t.Errorf("Expected PasswordHash %s, got %s", PasswordHash, req.PasswordHash)
	}
	if req.CreatedAt != CreatedAt {
		t.Errorf("Expected CreatedAt %d, got %d", CreatedAt, req.CreatedAt)
	}
	if req.IsBanned != IsBanned {
		t.Errorf("Expected IsBanned %d, got %d", IsBanned, req.IsBanned)
	}
	if req.IsSystem != IsSystem {
		t.Errorf("Expected IsSystem %d, got %d", IsSystem, req.IsSystem)
	}
	if req.UpdatedAt != UpdatedAt {
		t.Errorf("Expected UpdatedAt %d, got %d", UpdatedAt, req.UpdatedAt)
	}		
}
