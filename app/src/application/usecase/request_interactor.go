package usecase

import (
	"app/application/port"
	"app/domain"
	"app/messages"
	"log"
	"net/http"
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
func (i *SendFriendRequestInteractor) Execute(input SendFriendRequestInput) (int,error) {
	// 送信者とターゲットが同一でないか確認
	if input.SenderID == input.TargetID {
		return http.StatusBadRequest, messages.ErrSelfRequest
	}

	// 送信者とターゲットが存在するか確認
	senderExists, err := i.UserRepo.ExistsByID(input.SenderID)
	if err != nil {
		log.Println("Error checking sender existence:", err)
		return http.StatusInternalServerError, err
	}
	if !senderExists {
		return http.StatusNotFound, messages.ErrSenderNotFound
	}
	targetExists, err := i.UserRepo.ExistsByID(input.TargetID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if !targetExists {
		return http.StatusNotFound, messages.ErrTargetNotFound
	}

	// 既にリクエスト済みか確認
	exists, err := i.Repo.Exists(input.SenderID, input.TargetID)
	if err != nil {
		return http.StatusInternalServerError, err
	}	
	if exists {
		return http.StatusConflict, messages.ErrAlreadySent
	}

	// uuidを生成
	uid, err := i.Genid.Genid()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// ドメインエンティティの作成
	req, err := domain.SendFriendRequest(input.SenderID, input.TargetID, uid)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// 永続化
	return http.StatusCreated, i.Repo.Create(req)
}
