package main

import (
	"fmt"
	negotiate_response "negotiate-response"
	"net/http"
)

const PORT = ":7166"

func main() {
	http.HandleFunc("/", negotiate_response.Handler)

	fmt.Printf("Listening on %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
