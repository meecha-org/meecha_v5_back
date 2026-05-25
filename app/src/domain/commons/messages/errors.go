package commons

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

type ErrorType int

const (
	TypeNotFound ErrorType = iota
	TypeValidation
	TypeConflict
	TypeUnAuthorized
	TypeForbidden
	TypeBadRequest
	TypeInternal
)

func (t ErrorType) String() string {
	switch t {
	case TypeNotFound:
		return "NotFound"
	case TypeValidation:
		return "Validation"
	case TypeConflict:
		return "Conflict"
	case TypeUnAuthorized:
		return "UnAuthorized"
	case TypeForbidden:
		return "Forbidden"
	case TypeBadRequest:
		return "BadRequest"
	case TypeInternal:
		return "Internal"
	}

	return "Unknown"
}

type CustomError struct {
	Type    ErrorType
	Message string
	Cause   error
}

func NewNotFoundError(message string) error {
	return &CustomError{
		Type:    TypeNotFound,
		Message: message,
		Cause:   nil,
	}
}

func NewValidationError(message string) error {
	return &CustomError{
		Type:    TypeValidation,
		Message: message,
		Cause:   nil,
	}
}

func NewConflictError(message string) error {
	return &CustomError{
		Type:    TypeConflict,
		Message: message,
		Cause:   nil,
	}
}


func NewBadRequestError(message string) error {
	return &CustomError{
		Type:    TypeBadRequest,
		Message: message,
		Cause:   nil,
	}
}

func NewForbiddenError(message string) error {
	return &CustomError{
		Type:    TypeForbidden,
		Message: message,
		Cause:   nil,
	}
}

func (e *CustomError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}

	return e.Message
}
