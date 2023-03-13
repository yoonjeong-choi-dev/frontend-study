package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func longRunningProcess(w *io.PipeWriter) {
	// 10초간 느리게 요청 body 스트리밍
	for i := 0; i <= 10; i++ {
		log.Printf("Attack - %d\n", i)
		fmt.Fprintf(w, "attack!!")
		time.Sleep(1 * time.Second)
	}
	w.Close()
}

func main() {
	client := http.Client{}
	logReader, logWriter := io.Pipe()

	go longRunningProcess(logWriter)

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"http://localhost:7166/api",
		logReader,
	)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	log.Println("Starting client request")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error when reqeusting: %v\n", err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error when reading response : %v\n", err)
	}
	log.Printf("Response: %s\n", string(data))
}
