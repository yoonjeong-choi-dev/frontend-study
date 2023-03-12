package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Start processing the request")
	time.Sleep(15 * time.Second)
	fmt.Fprintf(w, "15 seconds passed...")
	log.Println("Finished processing the request")
	log.Println("This handler must be end despite of timeout handler middleware")
}

func handleUserAPIImproved(w http.ResponseWriter, r *http.Request) {
	log.Println("Start processing the request")
	time.Sleep(15 * time.Second)

	log.Println("Before continuing, it will check the timeout has already expired")
	if r.Context().Err() != nil {
		log.Printf("Aborting further processing: %v\n",
			r.Context().Err())

		// stop the handler when the connection with client is disconnected
		return
	}

	fmt.Fprintf(w, "15 seconds passed...")
	log.Println("Finished processing the request")
	log.Println("This handler cannot be reach here with timout handler middleware")
}

func main() {
	timeoutDuration := 5 * time.Second

	userHandler := http.HandlerFunc(handleUserAPI)
	timeoutHandler := http.TimeoutHandler(userHandler, timeoutDuration, "Out of Time...")

	userImprovedHandler := http.HandlerFunc(handleUserAPIImproved)
	timeoutImprovedHandler := http.TimeoutHandler(userImprovedHandler, timeoutDuration, "Out of Time...")

	mux := http.NewServeMux()
	mux.Handle("/simple", timeoutHandler)
	mux.Handle("/improve", timeoutImprovedHandler)

	log.Fatal(http.ListenAndServe(":7166", mux))
}
