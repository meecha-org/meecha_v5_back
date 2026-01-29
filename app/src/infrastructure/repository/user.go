package repository

import (
	"app/domain"
	"gorm.io/gorm"
)

// UserRepositoryImpl はGORMを使用したリポジトリ実装
type UserRepositoryImpl struct {
	DB *gorm.DB
}

// ExistsByID は指定されたユーザーIDが存在するか確認します
func (r *UserRepositoryImpl) ExistsByID(userID string) (bool, error) {
	var count int64
	err := r.DB.Model(&domain.User{}).
		Where(&domain.User{UserID: userID}).
		Count(&count).Error
	if err != nil {
		return false, err
	}	
	return count > 0, nil
}
