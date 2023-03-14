package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"

	users "user-service/service"
)

type userService struct {
	users.UnimplementedUsersServer
}

func (s *userService) GetUser(
	ctx context.Context, in *users.UserGetRequest) (*users.UserGetResponse, error) {
	log.Printf("Recieved request for user with Email: %s, Id: %s\n",
		in.Email, in.Id,
	)

	components := strings.Split(in.Email, "@")
	if len(components) != 2 {
		return nil, status.Error(
			codes.InvalidArgument, "invalid email address",
		)
	}

	user := users.User{
		Id:        in.Id,
		FirstName: components[0],
		LastName:  components[1],
		Age:       31,
	}
	return &users.UserGetResponse{User: &user}, nil
}

func registerService(s *grpc.Server) {
	users.RegisterUsersServer(s, &userService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	// TCP connection for gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Listener: %v\n", err)
	}

	server := grpc.NewServer()
	registerService(server)
	log.Fatal(startServer(server, listener))
}
