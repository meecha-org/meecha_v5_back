package domain

import (
	"app/domain/commons/messages"
	"time"
)

type FriendRequest struct {
	requestID string
	senderID  string
	targetID  string
	created   int64
}

// Getters
func (f *FriendRequest) RequestID() string { return f.requestID }
func (f *FriendRequest) SenderID() string  { return f.senderID }
func (f *FriendRequest) TargetID() string  { return f.targetID }
func (f *FriendRequest) Created() int64    { return f.created }

// 新規作成時：ビジネスルールを強制するコンストラクタ
func NewFriendRequest(requestID, senderID, targetID string) (*FriendRequest, error) {
	if senderID == targetID {
		return nil, commons.NewBadRequestError("ErrSelfRequest")
	}
	// IDの形式チェックなどもここに書く
	return &FriendRequest{
		requestID: requestID,
		senderID:  senderID,
		targetID:  targetID,
		created:   time.Now().Unix(),
	}, nil
}

// 再構築時：DBからの復元用（検証不要な場合）
func ReconstructFriendRequest(requestID, senderID, targetID string, created int64) *FriendRequest {
	return &FriendRequest{
		requestID: requestID,
		senderID:  senderID,
		targetID:  targetID,
		created:   created,
	}
}

// ドメインロジック：リクエストの有効期限を判定する例
func (f *FriendRequest) IsExpired(duration time.Duration) bool {
	return time.Now().Unix()-f.created > int64(duration.Seconds())
}
