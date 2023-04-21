package main

import (
	"fmt"
	"google.golang.org/grpc"
	"greeter"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	greeter.RegisterGreeterServiceServer(grpcServer, &Greeter{Exclaim: true})

	listener, err := net.Listen("tcp", ":7166")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listen on port 7166")
	fmt.Println(grpcServer.Serve(listener))
}
