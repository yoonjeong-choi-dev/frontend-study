package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// Wrapping grpc Client Stream for Logging Streaming
type wrappedClientStream struct {
	grpc.ClientStream
	msgSendCount     int
	msgReceivedCount int
}
type streamDurationContextKey struct{}
type streamDurationContextValue struct {
	startTime time.Time
}

func (s wrappedClientStream) SendMsg(m interface{}) error {
	log.Printf("Send Message: %T\n", m)
	err := s.ClientStream.SendMsg(m)
	s.msgSendCount += 1
	return err
}

func (s *wrappedClientStream) RecvMsg(m interface{}) error {
	log.Printf("Received Message: %T\n", m)
	err := s.ClientStream.RecvMsg(m)
	s.msgReceivedCount += 1
	return err
}

func (s *wrappedClientStream) CloseSend() error {
	log.Println("CloseSend() called")
	v := s.Context().Value(streamDurationContextKey{})

	if m, ok := v.(streamDurationContextValue); ok {
		log.Printf("Duration:%v", time.Since(m.startTime))
	}
	err := s.ClientStream.CloseSend()
	log.Printf("Messages Sent Number: %d, Messages Received Number: %d\n",
		s.msgSendCount,
		s.msgReceivedCount,
	)
	return err
}

func LoggingUnaryInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	conn *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, conn, opts...)
	end := time.Now()

	log.Printf("Method:%s, Duration:%s, Error:%v\n", method, end.Sub(start), err)

	return err
}

func LoggingStreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	conn *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	durationCtxValue := streamDurationContextValue{startTime: time.Now()}
	ctxWithTimer := context.WithValue(
		ctx,
		streamDurationContextKey{}, durationCtxValue)

	stream, err := streamer(ctxWithTimer, desc, conn, method, opts...)
	clientStream := &wrappedClientStream{
		ClientStream:     stream,
		msgReceivedCount: 0,
		msgSendCount:     0,
	}
	return clientStream, err
}
