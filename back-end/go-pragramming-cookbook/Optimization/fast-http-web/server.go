package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	controller := NewItemController()

	router := fasthttprouter.New()
	router.GET("/item", controller.GetItemsHandler)
	router.POST("/item/:item", controller.AddItemsHandler)

	fmt.Println(fasthttp.ListenAndServe(":7166", router.Handler))
}
