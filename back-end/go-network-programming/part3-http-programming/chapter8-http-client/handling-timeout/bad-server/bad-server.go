package main

import (
	"fmt"
	"net/http"
)

func badHandler(w http.ResponseWriter, r *http.Request) {
	// Blocking -> 요청을 무한정 대기하는 핸들러
	select {}
}

func main() {
	http.HandleFunc("/", badHandler)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
