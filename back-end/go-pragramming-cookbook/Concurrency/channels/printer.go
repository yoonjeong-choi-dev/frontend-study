package channels

import (
	"context"
	"fmt"
	"time"
)

// Printer 데이터 채널 소비자
func Printer(ctx context.Context, ch chan string) {
	t := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Printer end with context done")
			return
		case res := <-ch:
			fmt.Printf("From channel: %s\n", res)
		case <-t:
			fmt.Println("tok")
		}
	}
}
