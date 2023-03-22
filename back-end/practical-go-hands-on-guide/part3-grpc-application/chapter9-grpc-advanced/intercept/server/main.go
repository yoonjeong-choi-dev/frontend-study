package main

import (
	"google.golang.org/grpc"
	svc "intercept/service"
	"intercept/utils"
	"log"
	"net"
)

func createServerWithInterceptor() *grpc.Server {
	return grpc.NewServer(
		grpc.UnaryInterceptor(loggingUnaryInterceptor),
		grpc.StreamInterceptor(loggingStreamInterceptor),
	)
}

func registerService(s *grpc.Server) {
	svc.RegisterUsersServer(s, &userService{})
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Listener: %s\n",
			utils.GetJsonStringUnsafe(err),
		)
	}

	s := createServerWithInterceptor()
	registerService(s)
	log.Fatal(startServer(s, l))
}
