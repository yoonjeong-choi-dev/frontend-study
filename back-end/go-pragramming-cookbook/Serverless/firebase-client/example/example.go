package main

import (
	"context"
	"firebase"
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %s\n", r)
			fmt.Printf("stacktrace:\n%s\n", debug.Stack())
		}

	}()

	ctx := context.Background()
	client, err := firebase.Authenticate(ctx, "yj-collection")
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	if err := client.Set(ctx, "yj-key", []string{"yoonjeong", "choi", "dev"}); err != nil {
		panic(err)
	}

	res, err := client.Get(ctx, "yj-key")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stored Data: %#v\n", res)

	if err := client.Set(ctx, "another-yj-key", []string{"yj", "7166"}); err != nil {
		panic(err)
	}

	res, err = client.Get(ctx, "another-yj-key")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Stored Data: %#v\n", res)
}
