package domain

import (
	"app/domain/commons/messages"
	"time"
)

// FriendRequest はフレンドリクエストを表すドメインエンティティ
type FriendRequest struct {
	RequestID string // リクエストの一意なID (UUIDなど)
	SenderID  string // 送信者のユーザーID
	TargetID  string // ターゲットのユーザーID
	Created   int64  // 作成時間 (UNIX秒)
}

// NewFriendRequest は新しいフレンドリクエストを作成します。
func SendFriendRequest(senderID, targetID, requestID string) (*FriendRequest, error) {
	// 送信者とターゲットが同一でないか確認
	if senderID == targetID {
		return nil,commons.NewBadRequestError("ErrSelfRequest") // 400 Bad Request
	}
	return &FriendRequest{
		RequestID: requestID, // ID生成はユースケース層で行う
		SenderID:  senderID,
		TargetID:  targetID,
		Created:   time.Now().Unix(),
	}, nil
}
