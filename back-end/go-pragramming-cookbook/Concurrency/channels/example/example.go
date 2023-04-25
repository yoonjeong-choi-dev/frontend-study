package main

import (
	"channels"
	"context"
	"time"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go channels.Sender(ch, done)
	go channels.Printer(ctx, ch)

	// 2초 뒤에 Sender 고루틴 종료 신호 전송
	time.Sleep(2 * time.Second)
	done <- true

	// Printer 고루틴 종료를 위해 컨텍스트 종료
	cancel()

	// 고루틴들 정리를 위한 시간
	time.Sleep(3 * time.Second)
}
