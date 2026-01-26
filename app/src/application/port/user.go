package port

import (
	
) 

// UserPort はユーザーの永続化を行うインターフェース
type UserPort interface {
	// IDでユーザーが存在するか確認
	ExistsByID(userID string) (bool, error)
}
