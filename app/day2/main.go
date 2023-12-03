package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponseData struct {
	Status int 	   `json:"status"`
	Message string `json:"message"`
}

type RequestData struct {
    Message string `json:"message"`
}


func main() {
	e := echo.New()

	e.Use(middleware.Logger()) // ログ出力
	e.Use(middleware.Recover()) // 復帰用

	e.GET("/", hello)
	e.POST("/say", say)

	e.POST("error", func(c echo.Context) error {
		panic("error")
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	data := ResponseData{
		Status: 200,
		Message: "Hello, World!",
	}
	return c.JSON(http.StatusOK, data)
}

func say(c echo.Context) error {
	data := new(RequestData)
	if err := c.Bind(&data); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, data)
}