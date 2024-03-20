package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

func main() {
  // インスタンスを作成
  e := echo.New()

  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORS())
  // ルートを設定
  e.GET("/", hello) // ローカル環境の場合、http://127.0.0.2:8080/ にGETアクセスされるとhelloハンドラーを実行する

  // サーバーをポート番号8080で起動
  e.Logger.Fatal(e.Start(":8080"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	u := &User{
    Name:  "Jon",
    Email: "jon@labstack.com",
  }
  return c.JSON(http.StatusOK, u)
}