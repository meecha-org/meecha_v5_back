package repository

import (
	"app/domain"
	"gorm.io/gorm"
)

// DB専用の永続化モデル（インフラ層の都合）
type friendRequestModel struct {
	RequestID string `gorm:"primaryKey"`
	SenderID  string `gorm:"index"`
	TargetID  string `gorm:"index"`
	Created   int64
}

// GORMにテーブル名を指定
func (friendRequestModel) TableName() string {
	return "friend_requests"
}

type FriendRequestRepositoryImpl struct {
	DB *gorm.DB
}

// Create はドメインモデルを変換して保存
func (r *FriendRequestRepositoryImpl) Create(req *domain.FriendRequest) error {
	model := friendRequestModel{
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
	err := r.DB.Model(&friendRequestModel{}).
		Where("sender_id = ? AND target_id = ?", senderID, targetID).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
