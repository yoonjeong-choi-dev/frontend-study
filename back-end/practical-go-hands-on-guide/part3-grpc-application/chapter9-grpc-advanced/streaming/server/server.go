package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	svc "streaming/service"
)

func registerService(s *grpc.Server) {
	svc.RegisterUsersServer(s, &userService{})
	svc.RegisterReposServer(s, &repoService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Listener: %v\n", err)
	}

	s := grpc.NewServer()
	registerService(s)
	log.Fatal(startServer(s, l))
}
