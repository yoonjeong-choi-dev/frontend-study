package main

import (
	"context"
	"fmt"
	"net/http"
	"service"
)

const grpcHost = "http://localhost:7166"

func main() {
	client := service.NewGreeterServiceProtobufClient(grpcHost, &http.Client{})

	ctx := context.Background()
	req := service.GreetRequest{
		Name:     "YJ Choi",
		Greeting: "Hello~~~",
	}

	res, err := client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response 1: %s\n", res)

	req.Greeting = "Goodbye..."
	res, err = client.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response 2: %s\n", res)
}
