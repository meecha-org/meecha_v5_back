package domain

import "time"


const (
	Google    ProviderCode = "google"
	Github    ProviderCode = "github"
	Discord   ProviderCode = "discord"
	Microsoft ProviderCode = "microsoftonline"
	Basic     ProviderCode = "basic"
)

type ProviderCode string

type User struct {
	UserID       string       // ユーザーID
	Name         string       // ユーザー名
	Email        string       // メールアドレス
	ProvCode     ProviderCode // 認証プロバイダコード
	ProvUID      string       // 認証プロバイダUID
	PasswordHash string       // ハッシュ化されたパスワード
	CreatedAt    int64        // ユーザー作成日 (UNIX秒)
	IsBanned     int          // 0: 通常, 1: バン
	IsSystem     int          // 0: 通常, 1: システムユーザー
	UpdatedAt    int64        // 最終更新日 (UNIX秒)
}

// NewUser は新しいユーザーエンティティを作成します。
func NewUser(userID, name, email, provUID, passwordHash string, provCode ProviderCode) *User {
	return &User{
		UserID:       userID,
		Name:         name,
		Email:        email,
		ProvCode:     provCode,
		ProvUID:      provUID,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().Unix(),
		IsBanned:     0,
		IsSystem:     0,
		UpdatedAt:    time.Now().Unix(),
	}
}
