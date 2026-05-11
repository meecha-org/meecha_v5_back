package models

type User struct {
	UserID       string       `gorm:"type:varchar(255);primaryKey"`                             // ユーザーID
	Name         string       `gorm:"type:varchar(255)"`                                        // ユーザー名
	Email        string       `gorm:"type:varchar(255);uniqueIndex:idx_users_email,length:255"` // メールアドレス
	ProvCode     string       `gorm:"type:varchar(255);index:idx_prov_code,length:255"`         // 認証プロバイダコード
	ProvUID      string       `gorm:"type:varchar(255);index:idx_prov_uid,length:255"`          // 認証プロバイダUID
	PasswordHash string       `gorm:"default:''"`                                               // ハッシュ化されたパスワード
	CreatedAt    int64        `gorm:"autoCreateTime"`                                           // ユーザー作成日
	Sessions     []Session    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`            // ユーザーが持つセッション
	IsBanned     int          `gorm:"default:0"`                                                // ユーザーの禁止状態
	IsSystem     int          `gorm:"default:0"`                                                // システムユーザーかどうか
	Labels       []Label      `gorm:"many2many:user_labels;constraint:OnDelete:CASCADE"`        // ユーザーのラベル
	UpdatedAt    int64        `gorm:"autoUpdateTime"`                                           // ユーザー更新日
}

type Session struct {
	SessionID string `gorm:"primaryKey"` // セッションID
	UserID    string // ユーザーID
	UserAgent string // ユーザーエージェント
	RemoteIP  string // リモートIP
	CreatedAt int64  `gorm:"autoCreateTime"` // セッション作成日
}

type Label struct {
	ID    uint   `gorm:"primarykey"`      // ラベルのプライマリキー
	Name  string `gorm:"unique;not null"` // ラベル名（ユニークかつNULL不可）
	Color string `gorm:"default:#000000"` // ラベルの色
	CreatedAt int64 `gorm:"autoCreateTime"` // ラベルの作成日時
	// これも同じ中間テーブル "user_labels" を指定します。
	Users []*User `gorm:"many2many:user_labels;constraint:OnDelete:CASCADE"`
}


