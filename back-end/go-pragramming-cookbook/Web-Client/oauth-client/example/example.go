package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	oauth_client "oauth-client"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Printf("Stack Trace:\n%s\n", debug.Stack())
		}
	}()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := oauth_client.Setup()
	ctx := context.Background()

	token, err := oauth_client.GetToken(ctx, config)
	if err != nil {
		panic(err)
	}

	client := config.Client(ctx, token)
	if err := oauth_client.GetUserInfo(client); err != nil {
		panic(err)
	}
}
