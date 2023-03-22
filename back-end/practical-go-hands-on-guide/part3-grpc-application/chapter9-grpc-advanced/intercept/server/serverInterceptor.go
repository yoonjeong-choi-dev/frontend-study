package main

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
	var requestId string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No Meta Data")
	} else {
		if len(md.Get("Request-Id")) != 0 {
			requestId = md.Get("Request-Id")[0]
		}
	}

	log.Printf("Method:%s, Duration:%s, Error:%v, Request-Id:%s",
		method,
		latency,
		err,
		requestId,
	)
}

func loggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	logMessage(ctx, info.FullMethod, time.Since(start), err)
	return resp, err
}

func loggingStreamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	start := time.Now()
	err := handler(srv, stream)
	ctx := stream.Context()
	logMessage(ctx, info.FullMethod, time.Since(start), err)
	return err
}
