package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.TimeoutHandler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

				log.Println("Now we sleep with 10s")
				time.Sleep(10 * time.Second)
				log.Println("Woke up. Now we write the response")
			}),
		time.Second,
		"Timeout while reading response",
	)

	http.Handle("/", handler)
	fmt.Println(http.ListenAndServe(":7166", nil))
}
