package presentation

import (
	"app/application/usecase"
	"app/domain/commons/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

// FriendHandler フレンド関連のハンドラー
type FriendHandler struct {
	Interactor *usecase.SendFriendRequestInteractor //  Interactorへの依存を追加
}

// NewFriendHandler FriendHandlerのコンストラクタ
func NewFriendHandler(interactor *usecase.SendFriendRequestInteractor) *FriendHandler{
	return &FriendHandler{
        Interactor: interactor,
    }
}

// HandleSendRequest 
func (h *FriendHandler) HandleSendRequest(c echo.Context) error {
	var input usecase.SendFriendRequestInput
    
    // 1. 入力データの取得とバリデーション (Echoの機能を利用)
    // JSONリクエストボディを直接 input 構造体にバインド
	if err := c.Bind(&input); err != nil {
		// バインドエラーは通常 400 Bad Request
        // c.Logger.Error(err) 
		return c.JSON(http.StatusBadRequest, map[string]string{"error": messages.ErrInvalidFormat.Error()})
	}

    // 2. ユースケースの実行 (内側の層への呼び出し)
	code, err := h.Interactor.Execute(input)

	// 3. 結果の判定とレスポンスの返却
	if err != nil {
        // その他のエラーはサーバーエラーとして処理
        // c.Logger.Error(err) 
		return c.JSON(code, map[string]string{"error":err.Error()}) // 500 Internal Server Error
	}

    // 4. 成功レスポンス
    // 成功したリソース作成は 201 Created で返す
	return c.JSON(http.StatusCreated, map[string]string{"message": messages.SuccessRequestSent}) 
}
