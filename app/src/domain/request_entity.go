package domain

import (
	"time"
	"errors"
)

// FriendRequest はフレンドリクエストを表すドメインエンティティ
type FriendRequest struct {
	SenderID  string `gorm:"primaryKey"`
	TargetID  string `gorm:"primaryKey"`
	RequestID string `gorm:"size:36;uniqueIndex"` // UUIDなどを格納
	Created   int64  `gorm:"autoCreateTime"`      // 作成時間 (UNIX秒)
}

// NewFriendRequest は新しいフレンドリクエストを作成します。
func SendFriendRequest(senderID, targetID, requestID string) (*FriendRequest, error) {
	// 送信者とターゲットが同一でないか確認
	if senderID == targetID {
		return nil, errors.New("cannot send friend request to oneself")
	}
	return &FriendRequest{
		SenderID:  senderID,
		TargetID:  targetID,
		RequestID: requestID, // ID生成はユースケース層で行う
		Created:   time.Now().Unix(),
	},nil
}
