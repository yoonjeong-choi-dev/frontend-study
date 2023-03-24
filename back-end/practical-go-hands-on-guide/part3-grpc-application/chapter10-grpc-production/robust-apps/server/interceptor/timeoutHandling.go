package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func TimeoutUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	var res interface{}
	var err error

	ctxWithTimeout, cancel := context.WithTimeout(
		ctx,
		300*time.Millisecond,
	)
	defer cancel()

	ch := make(chan error)
	go func() {
		res, err = handler(ctxWithTimeout, req)
		ch <- err
	}()

	select {
	case <-ctxWithTimeout.Done():
		// 타임아웃 발생
		log.Printf("%s: Timeout for Unary Interceptor\n", info.FullMethod)
		cancel()
		err = status.Error(
			codes.DeadlineExceeded,
			fmt.Sprintf("%s: deadline exceeded", info.FullMethod),
		)
		return res, err
	case <-ch:
		// 핸들러 처리 완료 시 진입
	}
	return res, err
}

func TimeoutStreamInterceptor(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	serverStream := wrappedServerTimeoutStream{
		ServerStream:    stream,
		ReceivedTimeout: 10 * time.Second,
	}
	err := handler(server, serverStream)
	return err
}
