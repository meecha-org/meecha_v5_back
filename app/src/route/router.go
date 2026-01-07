package route

import (
	"app/presentation"
	"net/http"

	"github.com/labstack/echo/v4"
)	

func InitServer(friendHandler *presentation.FriendHandler) *echo.Echo {
	// サーバー作成
	server := echo.New()

	

	server.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	friendgroup := server.Group("/friend")
	{
		friendgroup.POST("/request", friendHandler.HandleSendRequest)
		
	}

	return server
}

