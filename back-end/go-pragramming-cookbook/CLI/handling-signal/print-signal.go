package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchAndPrintSignal 고루틴으로 실행시키는 시그널 리스너
func CatchAndPrintSignal(ch chan os.Signal, done chan bool) {
	// 시그널이 발생할 때 까지 블로킹
	sig := <-ch
	fmt.Println("\nReceived Signal:", sig)

	// 각 신호에 대한 핸들러
	switch sig {
	case syscall.SIGINT:
		fmt.Println("Handling SIGINT")
	case syscall.SIGTERM:
		fmt.Println("Handling SIGTERM")
	default:
		fmt.Println("Cannot handle this signal")
	}

	// 핸들링 후 종료
	done <- true
}

func main() {
	sig := make(chan os.Signal)
	done := make(chan bool)

	// 핸들링한 시그널 등록
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// go routine for listener
	go CatchAndPrintSignal(sig, done)

	// 핸들링이 완료될 때까지 블로킹
	// => done 채널로 인해, 중지/종료 신호가 와도 프로그램이 종료되지 않음
	<-done

	fmt.Println("Exit the program...")
}
