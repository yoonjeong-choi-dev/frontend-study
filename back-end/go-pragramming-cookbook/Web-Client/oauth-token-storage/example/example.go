package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	oauth_token_storage "oauth-token-storage"
	"os"
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

	conf := oauth_token_storage.Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Scopes:       []string{"repo", "user"},
			Endpoint:     github.Endpoint,
		},
		Storage: &oauth_token_storage.FileStorage{
			Path: "local-token.txt",
		},
	}

	ctx := context.Background()
	token, err := conf.GetToken(ctx)
	if err != nil {
		panic(err)
	}

	client := conf.Client(ctx, token)
	res, err := client.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status Code: %d\n", res.StatusCode)
	fmt.Println("Response Body:")

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var jsonDecoded map[string]interface{}
	err = json.Unmarshal(data, &jsonDecoded)
	if err != nil {
		panic(err)
	}

	parsed, err := getJsonString(jsonDecoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed)

	fmt.Println("Get token Test")
	token, err = conf.GetToken(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token: %s\n", token)
}

func getJsonString(v interface{}) (string, error) {
	vJson, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}
	return string(vJson), nil
}
