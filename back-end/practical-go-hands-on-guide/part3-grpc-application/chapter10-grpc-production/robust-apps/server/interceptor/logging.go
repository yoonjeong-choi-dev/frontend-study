package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func logMessage(
	ctx context.Context,
	method string,
	latency time.Duration,
	err error,
) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No Meta Data")
	}

	log.Printf("[Log Interceptor]Method:%s, Duration:%s, Error:%v, Request-Id:%s",
		method,
		latency,
		err,
		md.Get("Request-Id"),
	)
}

func LoggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	res, err := handler(ctx, req)
	logMessage(ctx, info.FullMethod, time.Since(start), err)
	return res, err
}

func LoggingStreamInterceptor(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	serverStream := wrappedServerLoggingStream{ServerStream: stream}

	start := time.Now()
	err := handler(server, serverStream)
	ctx := stream.Context()
	logMessage(ctx, info.FullMethod, time.Since(start), err)
	return err
}
