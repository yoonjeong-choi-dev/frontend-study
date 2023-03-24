package interceptor

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	svc "robust-app/service"
	"robust-app/utils"
	"testing"
	"time"
)

func TestDisconnectUnaryInterceptor(t *testing.T) {
	req := svc.UserGetRequest{}
	unaryInfo := &grpc.UnaryServerInfo{
		FullMethod: "Users.GetUser",
	}
	testUnaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		time.Sleep(300 * time.Millisecond)
		return svc.UserGetRequest{}, nil
	}
	incomingCtx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)
	defer cancel()

	_, err := DisconnectUnaryInterceptor(incomingCtx, &req, unaryInfo, testUnaryHandler)

	if err == nil {
		t.Fatalf("Expected non-nil err but got nil err\n")
	}

	expectedErr := status.Errorf(
		codes.Canceled,
		"Users.GetUser: Request Canceled",
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

type disconnectStream struct {
	grpc.ServerStream
	CancelFunc context.CancelFunc
}

func (s disconnectStream) Context() context.Context {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)
	s.CancelFunc = cancel
	return ctx
}

func TestDisconnectStreamInterceptor(t *testing.T) {
	testServer := "testServer"
	testStream := disconnectStream{}
	streamInfo := &grpc.StreamServerInfo{
		FullMethod:     "Users.GetUser",
		IsClientStream: true,
		IsServerStream: true,
	}
	testHandler := func(server interface{}, stream grpc.ServerStream) (err error) {
		time.Sleep(200 * time.Millisecond)
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

	err := DisconnectStreamInterceptor(testServer, testStream, streamInfo, testHandler)
	if err == nil {
		t.Fatalf("Expected non-nil err but got nil err\n")
	}

	expectedErr := status.Errorf(
		codes.Canceled,
		"Users.GetUser: Request Canceled while streaming",
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
