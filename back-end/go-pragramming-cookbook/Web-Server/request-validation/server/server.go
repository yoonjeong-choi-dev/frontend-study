package main

import (
	"fmt"
	"net/http"
	request_validation "request-validation"
)

const PORT = ":7166"

func main() {
	c := request_validation.New()
	http.HandleFunc("/", c.Handler)

	fmt.Printf("Listening on %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
