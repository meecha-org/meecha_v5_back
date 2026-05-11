package models

import "app/domain"

// FriendRequest はフレンドリクエストを表すドメインエンティティ
type FriendRequest struct {
	RequestID string `gorm:"size:36;uniqueIndex"` // UUIDなどを格納
	SenderID  string `gorm:"primaryKey"`
	TargetID  string `gorm:"primaryKey"`
	Created   int64  `gorm:"autoCreateTime"`      // 作成時間 (UNIX秒)
}

func (fre *FriendRequest) ToDomain() *domain.FriendRequest {
	return &domain.FriendRequest{
		RequestID: fre.RequestID,
		SenderID:  fre.SenderID,
		TargetID:  fre.TargetID,
		Created:   fre.Created,
	}
}
