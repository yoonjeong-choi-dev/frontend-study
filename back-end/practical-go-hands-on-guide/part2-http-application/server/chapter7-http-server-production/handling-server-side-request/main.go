package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const PORT = "7166"

// Mock Server Side Request
func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling Ping!")
	fmt.Fprintln(w, "pong")
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Processing the Request")
	doSomethingInProcess()

	// 서버 작업 이후, 다른 서버로 네트워크 요청을 위한 준비
	req, err := http.NewRequestWithContext(
		r.Context(),
		"GET",
		fmt.Sprintf("http://localhost:%s/ping", PORT),
		nil,
	)
	if err != nil {
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	client := &http.Client{}
	log.Println("Outgoing HTTP Request Start...")

	// 다른 서버로 네트워크 요청
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error in outgoing request: %v\n", err)
		http.Error(
			w, err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Fprint(w, string(data))
	log.Println("Finished Processing the Request")
}

func doSomethingInProcess() {
	log.Println("Start - doSomethingInProcess")
	time.Sleep(2 * time.Second)
	log.Println("End - doSomethingInProcess")
}

func main() {
	timeoutDuration := 1 * time.Second
	userHandler := http.HandlerFunc(handleUserAPI)
	timeoutHandler := http.TimeoutHandler(userHandler, timeoutDuration, "Out of Time...")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/api", timeoutHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}
