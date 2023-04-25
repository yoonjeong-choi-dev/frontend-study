package pipeline_with_worker_pool

import (
	"context"
	"fmt"
)

// Print in 채널의 데이터를 출력 후, out 채널로 전달
func (w *Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			fmt.Printf("[Printer %d] %s\n", w.id, val)
			w.out <- val
		}
	}
}
