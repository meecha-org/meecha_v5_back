package messages

import "errors"

// エラーメッセージの定義(変更不可)
var (
	ErrInvalidFormat  = errors.New("フォーマットが不正です")
	ErrAlreadySent    = errors.New("既にフレンドリクエストが送信されています")
	ErrSelfRequest    = errors.New("自分自身に送信することはできません")
	ErrSenderNotFound = errors.New("送信者が存在しません")
	ErrTargetNotFound = errors.New("ターゲットが存在しません")
)

//成功
const (
	SuccessRequestSent = "フレンドリクエストが送信されました"
)	
