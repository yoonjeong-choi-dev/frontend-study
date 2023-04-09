package main

import (
	"context"
	"io"
	"log"
	"ping"
	"time"
)

func main() {
	// mock connection by in-memory buffer
	r, w := io.Pipe()

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second

	// ping 요청을 위한 고루틴
	go func() {
		ping.PingWithInterval(ctx, w, resetTimer)
		close(done)
	}()

	receivePingHandler := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			log.Printf("Reset timer as %s by sending reset chan\n", d)
			resetTimer <- d
		}

		now := time.Now()

		// ping 에 대한 응답 받기
		buf := make([]byte, 1024)
		size, err := r.Read(buf)
		if err != nil {
			log.Printf("Error for reading response: %s\n", err.Error())
		} else {
			log.Printf("Recived %s by %s\n",
				string(buf[:size]),
				time.Since(now).Round(100*time.Millisecond),
			)
		}
	}

	durations := []int64{0, 200, 300, 0, -1, -1, -1}
	for idx, duration := range durations {
		log.Printf("Ping %d:\n", idx+1)
		receivePingHandler(time.Duration(duration)*time.Millisecond, r)
	}

	cancel()
	<-done
}
