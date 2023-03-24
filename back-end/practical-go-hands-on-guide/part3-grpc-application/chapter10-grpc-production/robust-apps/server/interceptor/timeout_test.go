package interceptor

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	svc "robust-app/service"
	"robust-app/utils"
	"testing"
	"time"
)

func TestTimeoutUnaryInterceptor(t *testing.T) {
	req := svc.UserGetRequest{}
	unaryInfo := &grpc.UnaryServerInfo{
		FullMethod: "Users.GetUser",
	}
	testUnaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		time.Sleep(500 * time.Millisecond)
		return svc.UserGetRequest{}, nil
	}

	_, err := TimeoutUnaryInterceptor(
		context.Background(),
		&req,
		unaryInfo,
		testUnaryHandler,
	)

	if err == nil {
		t.Fatalf("Expected non-nil err but got nil err\n")
	}

	expectedErr := status.Errorf(
		codes.DeadlineExceeded,
		"Users.GetUser: deadline exceeded",
	)
	if !errors.Is(err, expectedErr) {
		t.Errorf(
			"Expected error: %s, but Got %s\n",
			utils.GetJsonStringUnsafe(expectedErr.Error()),
			utils.GetJsonStringUnsafe(err.Error()),
		)
	}
	t.Logf("Error: %s\n", utils.GetJsonStringUnsafe(err.Error()))
}

type timeoutStream struct {
	grpc.ServerStream
}

func (s timeoutStream) SendMsg(m interface{}) error {
	log.Println("Test Stream - Send")
	return nil
}

func (s timeoutStream) RecvMsg(m interface{}) error {
	log.Println("Test Stream - Receive")
	// Rise Server Timeout
	time.Sleep(1500 * time.Millisecond)
	return nil
}

func TestTimeoutStreamInterceptor(t *testing.T) {
	testServer := "testServer"
	testStream := timeoutStream{}
	streamInfo := &grpc.StreamServerInfo{
		FullMethod:     "Users.GetUser",
		IsClientStream: true,
		IsServerStream: true,
	}
	testHandler := func(server interface{}, stream grpc.ServerStream) (err error) {
		for {
			req := svc.UserHelpRequest{}
			err := stream.RecvMsg(&req)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			res := svc.UserHelpResponse{}
			err = stream.SendMsg(&res)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := TimeoutStreamInterceptor(testServer, testStream, streamInfo, testHandler)
	if err == nil {
		t.Fatalf("Expected non-nil err but got nil err\n")
	}

	expectedErr := status.Errorf(
		codes.DeadlineExceeded,
		"deadline exceeded while streaming",
	)
	if !errors.Is(err, expectedErr) {
		t.Errorf(
			"Expected error: %s, but Got %s\n",
			utils.GetJsonStringUnsafe(expectedErr.Error()),
			utils.GetJsonStringUnsafe(err.Error()),
		)
	}
	t.Logf("Error: %s\n", utils.GetJsonStringUnsafe(err.Error()))
}
