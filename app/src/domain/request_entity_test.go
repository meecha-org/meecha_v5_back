package domain_test

import (
	"testing"
	"time"
    
	"app/domain"
)

// TestSendFriendRequest_Creation テスト: FriendRequestエンティティが正しく作成されるか
func TestSendFriendRequest_Creation(t *testing.T) {
	senderID := "user123"
	targetID := "user456"
	requestID := "uuid-test-123"
    
	// 実行前の現在時刻を取得（Createdフィールドのテスト用）
	startTime := time.Now().Unix()

	req := domain.SendFriendRequest(senderID, targetID, requestID)

	// 1. 各フィールドが正しく設定されているか確認
	if req.SenderID != senderID {
		t.Errorf("Expected SenderID %s, got %s", senderID, req.SenderID)
	}
	if req.TargetID != targetID {
		t.Errorf("Expected TargetID %s, got %s", targetID, req.TargetID)
	}
	if req.RequestID != requestID {
		t.Errorf("Expected RequestID %s, got %s", requestID, req.RequestID)
	}
    
	// 2. Created時間が正しく設定されているか（テスト実行開始後か）
	if req.Created < startTime || req.Created > time.Now().Unix() {
		t.Errorf("Created time is outside expected range: %d", req.Created)
	}
}

