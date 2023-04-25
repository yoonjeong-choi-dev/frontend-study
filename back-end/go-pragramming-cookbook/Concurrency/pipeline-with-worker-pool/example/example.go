package main

import (
	"context"
	"fmt"
	pipeline "pipeline-with-worker-pool"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// in -> out 중간의 파이프라인이 무엇을 하는지 클라이언트는 알지 못함
	// => 데이터 프로세싱 추상화
	in, out := pipeline.NewPipeline(ctx, 10, 2)

	numData := 20
	go func() {
		for i := 0; i < numData; i++ {
			in <- fmt.Sprintf("Data %d to process", i)
		}
	}()

	for i := 0; i < numData; i++ {
		<-out
	}
}
