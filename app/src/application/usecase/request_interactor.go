package usecase

import (
	"app/application/port"
	"app/domain"
	"app/messages"
	"log"
)

// SendFriendRequestInput はユースケースへの入力データ
type SendFriendRequestInput struct {
	SenderID string
	TargetID string
}

// SendFriendRequestInteractor はフレンドリクエスト送信のユースケース
type SendFriendRequestInteractor struct {
	Repo     port.FriendRequestPort
	Genid    port.UUIDGeneratorPort
	UserRepo port.UserPort
}

// Execute はユースケースを実行します
func (i *SendFriendRequestInteractor) Execute(input SendFriendRequestInput) error {
	// 送信者とターゲットが同一でないか確認
	if input.SenderID == input.TargetID {
		return messages.ErrSelfRequest
	}

	// 送信者とターゲットが存在するか確認
	senderExists, err := i.UserRepo.ExistsByID(input.SenderID)
	if err != nil {
		log.Println("Error checking sender existence:", err)
		return err
	}
	if !senderExists {
		return messages.ErrSenderNotFound
	}
	targetExists, err := i.UserRepo.ExistsByID(input.TargetID)
	if err != nil {
		return err
	}
	if !targetExists {
		return messages.ErrTargetNotFound
	}

	// 既にリクエスト済みか確認
	exists, err := i.Repo.Exists(input.SenderID, input.TargetID)
	if err != nil {
		return err
	}
	if exists {
		return messages.ErrAlreadySent
	}

	// uuidを生成
	uid, err := i.Genid.Genid()
	if err != nil {
		return err
	}

	// ドメインエンティティの作成
	req, err := domain.SendFriendRequest(input.SenderID, input.TargetID, uid)
	if err != nil {
		return err
	}

	// 永続化
	return i.Repo.Create(req)
}
