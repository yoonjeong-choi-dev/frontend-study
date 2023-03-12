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

	// 외부 요청이 오래 걸리는 상황
	time.Sleep(10 * time.Second)

	fmt.Fprintln(w, "pong")
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	log.Println("Start Processing the Request")

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
	log.Println("Got Response for Outgoing Http Request")

	// Outgoing 응답을 이용하여 요청 처리
	go func() {
		doSomethingInProcess(data)
		done <- true
	}()

	select {
	case <-done:
		log.Println("doSomethingInProcess done -> continuing request processing")
	case <-r.Context().Done():
		// 클라이언트 측에서 타임아웃 설정으로 먼저 요청을 끝는 경우
		// => 고루틴 실행 전에 끝낸 경우, 해당 고루틴(doSomethingInProcess) 실행 x
		log.Printf("Aborting request processing: %v\n", r.Context().Err())
		return
	}

	fmt.Fprint(w, string(data))
	log.Println("Finished Processing the Request")
}

func doSomethingInProcess(data []byte) {
	log.Println("Start - doSomethingInProcess")
	time.Sleep(15 * time.Second)
	log.Printf("Porcessed Done with %s\n", string(data))
	log.Println("End - doSomethingInProcess")
}

func main() {
	// 총 25초(ping-10s, doSomething-15s) 걸리므로, 서버 측에서는 타임아웃이 걸리지 않음
	timeoutDuration := 30 * time.Second
	userHandler := http.HandlerFunc(handleUserAPI)
	timeoutHandler := http.TimeoutHandler(userHandler, timeoutDuration, "Out of Time...")

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlePing)
	mux.Handle("/api", timeoutHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), mux))
}
