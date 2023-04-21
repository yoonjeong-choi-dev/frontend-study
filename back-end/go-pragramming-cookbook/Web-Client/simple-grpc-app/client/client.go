package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter"
)

func main() {
	conn, err := grpc.Dial(":7166", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()

	client := greeter.NewGreeterServiceClient(conn)
	ctx := context.Background()

	req := greeter.GreetRequest{
		Name:     "Yoonjeong",
		Greeting: "Hello~",
	}

	res, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response 1: %s\n", res)

	req.Greeting = "Goodbye~~~~"
	res, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response 2: %s\n", res)
}
