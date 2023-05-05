package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/guess", GuessHandler)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
