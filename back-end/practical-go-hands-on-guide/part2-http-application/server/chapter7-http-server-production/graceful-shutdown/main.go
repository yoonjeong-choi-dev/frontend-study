package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Start Processing the request")
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error when reading body: %v\n", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	log.Printf("Request Body: %s\n", string(data))
	fmt.Fprintf(w, "Response from API")
	log.Println("End Processing the request")
}

func shutDown(ctx context.Context, server *http.Server, waitForShutdownCompletion chan struct{}) {
	// signal.Notify 인자로 받은 시그널들이 발생할 때, 채널에 저장
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	log.Printf("Got signal: %#v. Server shutting down\n", sig)

	childCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := server.Shutdown(childCtx); err != nil {
		log.Printf("Error during shudonw: %v\n", err)
	}

	// main 고루틴에게 서버가 종료됨을 알려줌
	// => main 고루틴은 서버 종료 후 추가 작업을 한뒤 종료
	waitForShutdownCompletion <- struct{}{}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", handleAPI)

	server := http.Server{
		Addr:    ":7166",
		Handler: mux,
	}

	waitForShutdownCompletion := make(chan struct{})

	// 서버 종료 신호를 리스닝하는 고루틴 생성
	go shutDown(context.Background(), &server, waitForShutdownCompletion)

	err := server.ListenAndServe()
	log.Println("Wait for shutdown to complete")
	<-waitForShutdownCompletion
	log.Fatal(err)
}
