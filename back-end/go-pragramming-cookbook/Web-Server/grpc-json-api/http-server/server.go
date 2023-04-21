package main

import (
	"fmt"
	"grpc-json-api/http-server/handlers"
	"grpc-json-api/internal"
	"net/http"
)

const PORT = ":7166"

func main() {
	c := handlers.Controller{KeyValue: internal.NewKeyValue()}
	http.HandleFunc("/set", c.SetHandler)
	http.HandleFunc("/get", c.GetHandler)

	fmt.Printf("HTTP Server on port %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
