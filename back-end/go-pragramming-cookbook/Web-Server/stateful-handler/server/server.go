package main

import (
	"fmt"
	"net/http"
	stateful_handler "stateful-handler"
)

const PORT = ":7166"

func main() {
	storage := stateful_handler.InMemoryStorage{}
	controller := stateful_handler.NewController(&storage)

	// register handlers
	http.HandleFunc("/get", controller.GetValue(false))
	http.HandleFunc("/get/default", controller.GetValue(true))
	http.HandleFunc("/set", controller.SetValue)

	fmt.Printf("Listening on port %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
