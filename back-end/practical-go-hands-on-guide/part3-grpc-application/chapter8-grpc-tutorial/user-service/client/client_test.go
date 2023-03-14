package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	users "user-service/service"
)

type dummyUserService struct {
	users.UnimplementedUsersServer
}

var dummyUser = users.User{
	Id:        "test-user",
	FirstName: "first-name",
	LastName:  "last-name",
	Age:       31,
}

func (s *dummyUserService) GetUser(
	ctx context.Context, in *users.UserGetRequest) (*users.UserGetResponse, error) {
	return &users.UserGetResponse{User: &dummyUser}, nil
}

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	users.RegisterUsersServer(s, &dummyUserService{})

	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()
	return s, l
}

func TestGetUser(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	// Server Test 와 동일
	// grpc.WithContextDialer 매개변수로 이용하기 위한 래핑 함수
	// => net.Dial 통해 인메모리 통신
	bufconnDial := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}

	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(bufconnDial),
	)
	if err != nil {
		t.Fatalf("Error for creating grpc client: %v\n", err)
	}

	client := getUserServiceClient(conn)
	result, err := getUser(
		client,
		&users.UserGetRequest{
			Email: "test@test",
		},
	)
	if err != nil {
		t.Fatalf("Error for request: %v\n", err)
	}

	if result.User.FirstName != dummyUser.FirstName ||
		result.User.LastName != dummyUser.LastName {
		t.Errorf("Expected: %v\nGot: %v\n", dummyUser, result.User)
	}

	t.Logf("\nResponse: %v\n", result.User)
}
