package repository

import (
	"app/domain"
	"app/infrastructure/models"

	"gorm.io/gorm"
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
func (r *FriendRequestRepositoryImpl) Exists(senderID, targetID string) (bool, error) {
	var count int64
	// 検索時は構造体ではなく条件式を使うのがGoのORMにおける安全なプラクティスです
	err := r.DB.Model(&models.FriendRequest{}).
		Where("(sender_id = ? AND target_id = ?) OR (sender_id = ? AND target_id = ?)",
			senderID, targetID, targetID, senderID).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
