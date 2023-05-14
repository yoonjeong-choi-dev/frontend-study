package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 타임아웃 처리를 위한 컨텍스트
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 타임아웃이 설정된 컨텍스트를 통해 요청 생성
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:7166", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Timeout for request!")
		} else {
			panic(err)
		}
	} else {
		_ = res.Body.Close()
	}
}
