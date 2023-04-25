package pipeline_with_worker_pool

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
)

// Encode in 채널의 데이터 인코딩하여, out 채널로 전달
func (w *Worker) Encode(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-w.in:
			log.Printf("[Encoder %d] Encoded %s\n", w.id, val)
			w.out <- fmt.Sprintf("%s => %s",
				val,
				base64.StdEncoding.EncodeToString([]byte(val)),
			)
		}
	}
}
