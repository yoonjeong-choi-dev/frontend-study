package channels

import "time"

// Sender 100ms 단위로 "tick" 전송
func Sender(ch chan string, done chan bool) {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-done:
			// 완료 시, 해당 함수 종료
			ch <- "sender end"
			return
		case <-t:
			// 100ms 마다 채널에 전송
			ch <- "tick"
		}
	}
}
