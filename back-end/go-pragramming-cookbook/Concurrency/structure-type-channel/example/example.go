package main

import (
	"context"
	"fmt"
	ops "structure-type-channel"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reqChan := make(chan *ops.OpsRequest, 10)
	resChan := make(chan *ops.OpsResponse, 10)

	// 요청 처리를 위한 고루틴 실행
	go ops.Processor(ctx, reqChan, resChan)

	// 여러 연산 요청
	req1 := ops.OpsRequest{ops.Add, 1, 2}
	reqChan <- &req1

	req2 := ops.OpsRequest{ops.Subtract, 3, 4}
	reqChan <- &req2

	req3 := ops.OpsRequest{ops.Multiply, 5, 6}
	reqChan <- &req3

	req4 := ops.OpsRequest{ops.Divide, 7, 3}
	reqChan <- &req4

	req5 := ops.OpsRequest{ops.Divide, 1, 0}
	reqChan <- &req5

	// 응답 확인
	for i := 0; i < 5; i++ {
		res := <-resChan
		fmt.Printf("Request: %v, Result: %v, Error: %v\n",
			res.Request, res.Result, res.Err,
		)
	}
}
