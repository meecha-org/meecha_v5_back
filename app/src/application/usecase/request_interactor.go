package usecase

import (
	"app/application/port"
	"app/domain"
	"errors"
)

var (
	ErrAlreadySent = errors.New("friend request already sent")
)

// SendFriendRequestInput はユースケースへの入力データ
type SendFriendRequestInput struct {
	SenderID string
	TargetID string
}

// SendFriendRequestInteractor はフレンドリクエスト送信のユースケース
type SendFriendRequestInteractor struct {
	Repo  port.FriendRequestPort
	Genid port.UUIDGeneratorPort
}

// Execute はユースケースを実行します
func (i *SendFriendRequestInteractor) Execute(input SendFriendRequestInput) error {
	// 送信者とターゲットが同一でないか確認
	if input.SenderID == input.TargetID {
		return errors.New("cannot send friend request to oneself")
	}

	// 既にリクエスト済みか確認
	exists, err := i.Repo.Exists(input.SenderID, input.TargetID)
	if err != nil {
		return err
	}
	if exists {
		return ErrAlreadySent
	}

	// uuidを生成
	uid, err := i.Genid.Genid()
	if err != nil {
		return err
	}

	// ドメインエンティティの作成
	req := domain.SendFriendRequest(input.SenderID, input.TargetID, uid)

	// 永続化
	return i.Repo.Create(req)
}
