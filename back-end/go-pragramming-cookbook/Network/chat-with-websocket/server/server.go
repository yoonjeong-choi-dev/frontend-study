package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = "localhost:7166"

func main() {
	fmt.Printf("Listening on %s\n", addr)
	log.Panic(http.ListenAndServe(addr, http.HandlerFunc(websocketHandler)))
}
