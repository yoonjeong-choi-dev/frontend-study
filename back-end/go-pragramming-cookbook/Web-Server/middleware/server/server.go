package main

import (
	"fmt"
	"log"
	"middleware"
	"net/http"
	"os"
)

func main() {
	// 미들웨어 등록
	// 아래에 있는 미들웨어가 먼저 요청 처리
	handler := middleware.ApplyMiddlewares(
		middleware.GreetHandler,
		middleware.Logger(log.New(os.Stdout, "[middleware-logger]", 0)),
		middleware.RequestIdMiddleware(93),
	)

	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
