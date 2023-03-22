package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	svc "tls/service"
)

func SayHi(client svc.GreetingClient, username string) (*svc.GreetMessage, error) {
	req := svc.User{
		Name: username,
	}

	return client.SayHi(context.Background(), &req)
}

func main() {
	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if err != nil {
		log.Fatalf("Error for loaing cert files: %#v\n", err)
	}

	credOpts := grpc.WithTransportCredentials(creds)

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		credOpts,
		//grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("Error for grpc-connection: %#v\n", err)
	}

	client := svc.NewGreetingClient(conn)

	res, err := SayHi(client, "Yoonjeong")
	if err != nil {
		log.Printf("Error for request: %v\n", err)
	} else {
		log.Printf("Response: %s\n", res.Message)
	}

	res, err = SayHi(client, "")
	if err != nil {
		log.Printf("Error for request: %v\n", err)
	} else {
		log.Printf("Response: %s\n", res.Message)
	}
}
