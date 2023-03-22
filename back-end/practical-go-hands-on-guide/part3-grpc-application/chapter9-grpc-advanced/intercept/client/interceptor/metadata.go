package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// MetaDataUnaryInterceptor for unary pattern
func MetaDataUnaryInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	conn *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctxWithMetaData := metadata.AppendToOutgoingContext(
		ctx,
		"Request-Id",
		fmt.Sprintf("Client-Request-%s", method),
	)
	return invoker(ctxWithMetaData, method, req, reply, conn, opts...)
}

// MetaDataStreamInterceptor for client-side streaming
func MetaDataStreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	conn *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	ctxWithMetaData := metadata.AppendToOutgoingContext(
		ctx,
		"Request-Id",
		fmt.Sprintf("Client-Request-%s", method),
	)
	clientStream, err := streamer(
		ctxWithMetaData,
		desc,
		conn,
		method,
		opts...,
	)
	return clientStream, err
}
