package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong")
}

func handleSleep(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Processing the Request - handleSleep")

	query := r.URL.Query()
	duration := 3
	p := query.Get("time")
	v, err := strconv.Atoi(p)
	if err == nil {
		duration = v
	}

	log.Printf("Sleeping... with %d seconds\n", duration)
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Fprintf(w, "Done with time %d(s)", duration)

	log.Println("Finished Processing the Request - handleSleep")
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Processing the Request - handleAPI")
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v\n", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	log.Printf("Reqeust body: %s\n", string(data))
	fmt.Fprintf(w, "Response from API\n")

	log.Println("Finished Processing the Request - handleAPI")
}

func main() {
	timeoutForHandler := 5 * time.Second
	timeoutForReadRequest := 10 * time.Second   // handler 에게 전송할 요청 객체 생성 타임아웃
	timeoutForWriteResponse := 10 * time.Second // handler 에서 처리한 내용을 이용한 응답 객체 생성 타임아웃

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.HandleFunc("/sleep", handleSleep)
	mux.HandleFunc("/api", handleAPI)
	timeoutMux := http.TimeoutHandler(mux, timeoutForHandler, "Timeout for Handler")

	server := http.Server{
		Addr:         ":7166",
		Handler:      timeoutMux,
		ReadTimeout:  timeoutForReadRequest,
		WriteTimeout: timeoutForWriteResponse,
	}
	log.Fatal(server.ListenAndServe())
}
