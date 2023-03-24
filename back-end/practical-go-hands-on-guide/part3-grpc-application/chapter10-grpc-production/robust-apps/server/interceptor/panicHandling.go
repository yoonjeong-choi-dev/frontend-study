package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func PanicUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (res interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic Revocered: %v\n", r)
			err = status.Error(
				codes.Internal,
				"unexpected error occurred",
			)
		}
	}()

	res, err = handler(ctx, req)
	return
}

func PanicStreamInterceptor(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic Revocered: %v\n", r)
			err = status.Error(
				codes.Internal,
				"unexpected error occurred while streaming",
			)
		}
	}()
	serverStream := wrappedServerLoggingStream{ServerStream: stream}
	err = handler(server, serverStream)
	return
}
