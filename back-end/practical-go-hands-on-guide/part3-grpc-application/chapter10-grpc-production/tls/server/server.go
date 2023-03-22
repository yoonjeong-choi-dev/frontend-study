package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	svc "tls/service"
)

type greetingService struct {
	svc.UnimplementedGreetingServer
}

func (s *greetingService) SayHi(ctx context.Context, in *svc.User) (*svc.GreetMessage, error) {
	user := in.Name
	if len(user) == 0 {
		user = "anonymous"
	}

	res := svc.GreetMessage{
		Message: fmt.Sprintf("Hello~~ %s!", user),
	}
	return &res, nil
}

func registerService(s *grpc.Server) {
	svc.RegisterGreetingServer(s, &greetingService{})
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP: %#v\n", err)
	}

	creds, err := credentials.NewServerTLSFromFile(
		"../cert/server.crt",
		"../cert/server.key",
	)
	if err != nil {
		log.Fatalf("Error for loaing cert files: %#v\n", err)
	}
	credOpt := grpc.Creds(creds)

	server := grpc.NewServer(credOpt)
	registerService(server)
	log.Fatal(server.Serve(l))
}
