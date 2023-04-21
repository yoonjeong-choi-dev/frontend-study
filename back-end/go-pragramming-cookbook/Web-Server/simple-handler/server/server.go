package main

import (
	"fmt"
	"net/http"
	simple_handler "simple-handler"
)

const PORT = ":7166"

func main() {
	// register handler
	http.HandleFunc("/hello", simple_handler.HelloHandler)
	http.HandleFunc("/greet", simple_handler.GreetingHandler)

	fmt.Printf("Listening on port %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
