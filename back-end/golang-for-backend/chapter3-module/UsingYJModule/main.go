package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/yoonjeong-choi-dev/moloco-study/back-end/golang-for-backend/chapter3-module/YJModule"
	"net/http"
	"strconv"
)

const PORT = 7166

func main() {
	e := echo.New()

	// http://localhost:7166/
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello Go Webserver!")
	})

	// http://localhost:7166/{testNumber}
	e.GET("/:number", func(ctx echo.Context) error {
		numberParam := ctx.Param("number")
		number, err := strconv.Atoi(numberParam)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}

		isEven := YJModule.IsEven(number)
		isPositive := YJModule.IsPositive(number)
		return ctx.String(http.StatusOK, fmt.Sprintf("Your number is %d\n\tIs Positive?: %t\n\tIs Even?: %t\n", number, isPositive, isEven))
	})

	//
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
