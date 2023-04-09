package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	// create mock server
	server, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatalf("Error for server: %s\n", err.Error())
	}
	defer server.Close()

	// server handles only one connection
	go func() {
		conn, err := server.Accept()
		if err == nil {
			conn.Close()
		}
	}()

	// create context which can disconnect connection by cancel
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

	clientRequestDial := func(
		ctx context.Context,
		address string,
		response chan int,
		id int, wg *sync.WaitGroup) {

		defer wg.Done()

		var dialer net.Dialer

		// 동일한 context 객체를 이용하여 연결 생성
		// => 모든 클라이언트 측 연결은 하나의 context 공유
		conn, err := dialer.DialContext(ctx, "tcp", address)
		if err != nil {
			return
		}
		conn.Close()

		select {
		case <-ctx.Done():
			// 응답을 받지 못한 고루틴
		case response <- id:
			// 응답을 받은 고루틴
		}
	}

	// 총 10개의 요청을 동시에 생성
	// 공유 컨텍스트를 통해 응답을 받으면 전체 요청 취소
	res := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go clientRequestDial(ctx, server.Addr().String(), res, i, &wg)
	}

	// 하나의 응답을 받을 때까지 대기
	response := <-res

	// 응답 받은 후, 컨텍스트 취소를 통해 요청들에 대한 고루틴 종료
	cancel()
	wg.Wait()
	close(res)

	// result
	log.Printf("Dialer %d recieved the data\n", response)

	// error for context
	log.Printf("Canceled: %v\n", ctx.Err() == context.Canceled)
	log.Printf("DeadlineExceeded: %v\n", ctx.Err() == context.DeadlineExceeded)
	log.Printf("Context Error: %s\n", ctx.Err().Error())
}
