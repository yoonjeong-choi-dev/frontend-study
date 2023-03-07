package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func streamingHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})
	logReader, logWriter := io.Pipe()

	// 생산자 및 소비자 고루틴 생성
	go longRunningProcess(logWriter)
	go consumeStreamingProcess(logReader, w, done)

	<-done
}

// 생산자 : 스트리밍으로 처리할 데이터 입력
func longRunningProcess(logWriter *io.PipeWriter) {
	for i := 0; i <= 20; i++ {
		fmt.Fprintf(
			logWriter,
			`{"id": %d, "user_ip": "127.0. 0.1", "event":"click_the_create_button(%d)"`,
			i, i,
		)
		fmt.Fprintln(logWriter)

		// long process
		time.Sleep(1 * time.Second)
	}
	logWriter.Close()
}

// 소비자 : 생산자가 입력한 데이터 스트림을 읽어서 응답에 쓰는 역할
func consumeStreamingProcess(logReader *io.PipeReader, w http.ResponseWriter, done chan struct{}) {
	defer logReader.Close()

	buf := make([]byte, 500)
	flusher, flushSupported := w.(http.Flusher)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	// 생산자가 데이터를 모두 입력할 때까지 반복
	for {
		size, err := logReader.Read(buf)
		if err == io.EOF {
			break
		}

		// 응답에 해당 데이터 씀
		w.Write(buf[:size])
		// flush 가능한 경우, 바로 클라이언트에게 전송
		if flushSupported {
			flusher.Flush()
		}
	}

	// 모두 완료한 경우 다른 고루틴에게 알림
	done <- struct{}{}
}
