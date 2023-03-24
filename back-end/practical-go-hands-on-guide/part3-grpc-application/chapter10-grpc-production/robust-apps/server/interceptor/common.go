package interceptor

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type wrappedServerLoggingStream struct {
	grpc.ServerStream
}

func (s wrappedServerLoggingStream) SendMsg(m interface{}) error {
	log.Printf("[Log Wrapped Stream]Send msg called: %T\n", m)
	return s.ServerStream.SendMsg(m)
}

func (s wrappedServerLoggingStream) RecvMsg(m interface{}) error {
	log.Printf("[Log Wrapped Stream]Waiting to receive a message: %T\n", m)
	return s.ServerStream.RecvMsg(m)
}

type wrappedServerTimeoutStream struct {
	grpc.ServerStream
	ReceivedTimeout time.Duration
}

func (s wrappedServerTimeoutStream) SendMsg(m interface{}) error {
	return s.ServerStream.SendMsg(m)
}

func (s wrappedServerTimeoutStream) RecvMsg(m interface{}) error {
	ch := make(chan error)
	timer := time.NewTimer(s.ReceivedTimeout)
	go func() {
		log.Printf("[Timeout Wrapped Stream]Waiting to receive a message: %T\n", m)
		ch <- s.ServerStream.RecvMsg(m)
	}()

	select {
	case <-timer.C:
		log.Println("[Timeout Stream Interceptor]Timeout!")
		return status.Error(
			codes.DeadlineExceeded,
			fmt.Sprintf("deadline exceeded while streaming"),
		)
	case err := <-ch:
		return err
	}
}
