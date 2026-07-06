package commons

import "errors"

// NotFoundError はレコードが見つからない場合のエラーを表します
var NotFoundError = errors.New("Record not found") //"Record not found"
// InternalError は内部サーバーエラーを表します
var InternalError = errors.New("Internal server error") //NewInternalError
