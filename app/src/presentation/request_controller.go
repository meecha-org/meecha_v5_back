package presentation

import (
	"app/application/usecase"
	commons "app/domain/commons/messages"
	"app/presentation/utlis"
	"net/http"

	"github.com/labstack/echo/v4"
)

// FriendHandler フレンド関連のハンドラー
type FriendHandler struct {
	Interactor *usecase.SendFriendRequestInteractor //  Interactorへの依存を追加
}

// NewFriendHandler FriendHandlerのコンストラクタ
func NewFriendHandler(interactor *usecase.SendFriendRequestInteractor) *FriendHandler {
	return &FriendHandler{
		Interactor: interactor,
	}
}

// HandleSendRequest
func (h *FriendHandler) HandleSendRequest(c echo.Context) error {
	var input usecase.SendFriendRequestInput

	// JSONリクエストボディを直接 input 構造体にバインド
	if err := c.Bind(&input); err != nil {
		// c.Logger.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.NewBadRequestError("Invalid request").Error()}) // 400 Bad Request
	}

	//　ユースケースの実行 (内側の層への呼び出し)
	err := h.Interactor.Execute(input)

	utils.HandleError(c, err)

	// 成功したリソース作成は 201 Created で返す
	return c.JSON(http.StatusCreated, nil) // 201 Created
}
