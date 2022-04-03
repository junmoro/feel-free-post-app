package main

import (
	"net/http"

	"github.com/junmoro/feel-free-post-app/go-app/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)
	e.Logger.Fatal(e.Start(":8000"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()
	err := sqlDB.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "データベース接続失敗")
	} else {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
