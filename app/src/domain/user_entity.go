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
	UserID       string      
	Name         string       
	Email        string       
	ProvCode     ProviderCode
	ProvUID      string     
	PasswordHash string     
	CreatedAt    int64      
	IsBanned     int        
	IsSystem     int          
	UpdatedAt    int64
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
