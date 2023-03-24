package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func DisconnectUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	var res interface{}
	var err error

	ch := make(chan error)
	go func() {
		res, err = handler(ctx, req)
		ch <- err
	}()

	select {
	case <-ctx.Done():
		// disconnect
		err = status.Error(
			codes.Canceled,
			fmt.Sprintf("%s: Request Canceled", info.FullMethod),
		)
	case <-ch:
		// finish handling
	}
	return res, err
}

func DisconnectStreamInterceptor(
	server interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) (err error) {
	ch := make(chan error)
	go func() {
		err = handler(server, stream)
		ch <- err
	}()

	select {
	case <-stream.Context().Done():
		// disconnect
		err = status.Error(
			codes.Canceled,
			fmt.Sprintf("%s: Request Canceled while streaming", info.FullMethod),
		)
		return
	case <-ch:
		// finish handling
	}
	return
}
