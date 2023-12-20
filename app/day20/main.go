package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RequestData struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)

	e.POST("error", func(c echo.Context) error {
		panic("error")
	})
	e.Logger.Fatal(e.Start(":8081"))
}

func hello(c echo.Context) error {
	data := ResponseData{
		Status:  200,
		Message: "Hello by air",
	}
	return c.JSON(http.StatusOK, data)
}
