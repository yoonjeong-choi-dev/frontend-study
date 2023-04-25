package worker_pool_pattern

import (
	"context"
	"fmt"
)

func CryptoWorkerPool(ctx context.Context, numWorkers int) (context.CancelFunc, chan CryptoRequest, chan CryptoResponse) {
	// cancel 함수를 이용하여 워커 풀에 있는 워커들의 컨텍스트를 연결
	ctx, cancel := context.WithCancel(ctx)

	in := make(chan CryptoRequest, 10)
	out := make(chan CryptoResponse, 10)

	for i := 0; i < numWorkers; i++ {
		go worker(ctx, i, in, out)
	}

	return cancel, in, out
}

// worker 워커 풀에 생성되는 암호화 워커 프로세스
func worker(ctx context.Context, id int, in chan CryptoRequest, out chan CryptoResponse) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-in:
			fmt.Printf("[Worker %d] Processing %s\n", id, r.Op)
			out <- CryptoProcess(r)
		}
	}
}
