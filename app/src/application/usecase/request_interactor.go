package usecase

import (
	"app/application/port"
	"app/domain"
	commons "app/domain/commons/messages"
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
	// 送信者とターゲットが同一でないか確認(早期リターンでDB操作を減らす)
	if input.SenderID == input.TargetID {
		return commons.NewBadRequestError("ErrSelfRequest") // 400 Bad Request
	}

	// 送信者とターゲットが存在するか確認
	senderExists, err := i.UserRepo.ExistsByID(input.SenderID)
	if err != nil {
		return commons.NewInternalError(err.Error())
	}
	if !senderExists {
		return commons.NewBadRequestError("ErrSenderNotFound") // 400 Bad Request
	}
	targetExists, err := i.UserRepo.ExistsByID(input.TargetID)
	if err != nil {
		return commons.NewInternalError(err.Error())
	}
	if !targetExists {
		return commons.NewBadRequestError("ErrTargetNotFound") // 400 Bad Request
	}

	// 既にリクエスト済みか確認
	exists, serr := i.Repo.Exists(input.SenderID, input.TargetID)
	if serr != "" {
		return commons.NewInternalError(serr) //500 Internal Server Error
	}
	if exists {
		return commons.NewConflictError("ErrAlreadySent") // 409 Conflict
	}

	// uuidを生成
	uid, err := i.Genid.Genid()
	if err != nil {
		return commons.NewInternalError(err.Error())
	}

	// ドメインエンティティの作成
	req, err := domain.NewFriendRequest(uid, input.SenderID, input.TargetID)
	if err != nil {
		return commons.NewBadRequestError(err.Error()) // 400 Bad Request
	}

	// 永続化
	return i.Repo.Create(req)
}
