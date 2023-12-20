package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponseData struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger()) // ログ出力

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	data := ResponseData{
		Status: "ok",
	}
	return c.JSON(http.StatusOK, data)
}
