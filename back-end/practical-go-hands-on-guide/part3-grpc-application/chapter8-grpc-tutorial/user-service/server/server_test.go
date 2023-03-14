package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	users "user-service/service"
)

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	registerService(s)

	go func() {
		err := startServer(s, l)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()
	return s, l
}

func TestUserService_GetUser(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	// grpc.WithContextDialer 매개변수로 이용하기 위한 래핑 함수
	// => net.Dial 통해 인메모리 통신
	bufconnDialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}
	client, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatalf("Error for creating grpc client: %v\n", err)
	}

	userClient := users.NewUsersClient(client)

	// 인메모리 통신
	id := "mock-id"
	email := "yoonjeong@choi.com"
	resp, err := userClient.GetUser(
		context.Background(),
		&users.UserGetRequest{
			Id:    id,
			Email: email,
		},
	)

	if err != nil {
		t.Fatalf("Error for request: %v\n", err)
	}

	respEmail := fmt.Sprintf("%s@%s", resp.User.FirstName, resp.User.LastName)
	if resp.User.Id != id ||
		respEmail != email {
		t.Errorf("Expected Id to be %s, Got: %s\n", id, resp.User.Id)
		t.Errorf("Expected Email to be %s, Got: %s\n", email, respEmail)
	}

	t.Logf("\nResponse: %v\n", resp.User)
}
