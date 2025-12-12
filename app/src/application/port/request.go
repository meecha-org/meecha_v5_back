package port

import (
	"app/domain"
	) 

// FriendRequestPort はフレンドリクエストの永続化を行うインターフェース
type FriendRequestPort interface {
	// リクエストを作成
	Create(req *domain.FriendRequest) error 
	// 既にリクエストが存在するか確認
	Exists(senderID, targetID string) (bool, error)
}
