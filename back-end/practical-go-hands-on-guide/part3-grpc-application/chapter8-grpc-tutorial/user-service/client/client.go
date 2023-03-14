package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	users "user-service/service"
)

func createGrpcConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
}

func getUserServiceClient(conn *grpc.ClientConn) users.UsersClient {
	return users.NewUsersClient(conn)
}

func getUser(client users.UsersClient, u *users.UserGetRequest) (*users.UserGetResponse, error) {
	return client.GetUser(context.Background(), u)
}

func main() {
	conn, err := createGrpcConnection("localhost:50051")
	if err != nil {
		log.Fatalf("Error for grpc connection: %v\n", conn)
	}
	defer conn.Close()

	client := getUserServiceClient(conn)
	result, err := getUser(
		client,
		&users.UserGetRequest{Email: "yoonjeong@choi.com"},
	)
	if err != nil {
		log.Fatalf("Error calling users.GetUser: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "User: %s %s\n", result.User.FirstName, result.User.LastName)
}
