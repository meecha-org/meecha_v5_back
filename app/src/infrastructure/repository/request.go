package repository

import (
	"app/domain"
	"gorm.io/gorm"
)

// FriendRequestRepositoryImpl はGORMを使用したリポジトリ実装
type FriendRequestRepositoryImpl struct {
	DB *gorm.DB
}

// Create はフレンドリクエストをDBに登録します
func (r *FriendRequestRepositoryImpl) Create(req *domain.FriendRequest) error {
	// 構造体はdomain.FriendRequestをそのまま利用
	return r.DB.Create(req).Error
}

// Exists は既にリクエストが存在するか確認します
func (r *FriendRequestRepositoryImpl) Exists(senderID, targetID string) (bool, error) {
	var count int64
	err := r.DB.Model(&domain.FriendRequest{}).
		Where(&domain.FriendRequest{SenderID: senderID, TargetID: targetID}).
		Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
