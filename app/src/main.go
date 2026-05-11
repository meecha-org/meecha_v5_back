package main

import (
	"app/infrastructure/config"
	"app/route"
	
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// .env読み込み
	Init()

	//db接続、マイグレーション
	db := config.Init()

	//依存性解決とハンドラ取得
	friendHandler := route.SetupDependencies(db)

	// ルーティング設定
	server := route.InitServer(friendHandler)

	// ミドルウェア設定
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// サーバー起動 (EchoのListen)
	server.Logger.Fatal(server.Start(":9090")) // 💡 Echoの起動メソッドを使用
}
