package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-json-api/internal"
	"grpc-json-api/keyvalue"
	"net"
)

const PORT = ":4000"

func main() {
	grpcServer := grpc.NewServer()
	keyvalue.RegisterKeyValueServer(grpcServer, internal.NewKeyValue())

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Internel Server with port %s\n", PORT)
	fmt.Println(grpcServer.Serve(listener))
}
