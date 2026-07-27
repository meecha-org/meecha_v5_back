package repository

import (
	"app/domain"
	commons "app/domain/commons/messages"
	"app/infrastructure/models"

	"gorm.io/gorm"
	"app/domain/logger"
)

type FriendRequestRepositoryImpl struct {
	DB *gorm.DB
}

// Create はドメインモデルを変換して保存
func (r *FriendRequestRepositoryImpl) Create(req *domain.FriendRequest) error {
	model := models.FriendRequest{
		RequestID: req.RequestID(),
		SenderID:  req.SenderID(),
		TargetID:  req.TargetID(),
		Created:   req.Created(),
	}
	return r.DB.Create(&model).Error
}

// Exists はDBモデルを使用して存在確認
func (r *FriendRequestRepositoryImpl) Exists(senderID, targetID string) (bool, string) {
	var count int64
	// 検索時は構造体ではなく条件式を使うのがGoのORMにおける安全なプラクティスです
	err := r.DB.Model(&models.FriendRequest{}).
		Where("(sender_id = ? AND target_id = ?) OR (sender_id = ? AND target_id = ?)",
			senderID, targetID, targetID, senderID).Count(&count).Error

	if err == gorm.ErrRecordNotFound {
		logger.Println(err)
		return false, commons.NotFoundError.Error()
	}
	if err != nil {
		logger.Println(err)
		return false, commons.InternalError.Error()
	}

	return count > 0, ""
}
