package oauth_client

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"os"
)

func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

// GetToken 현재 설정을 이용하여 토큰을 받아올 url 가져온 뒤,
// 브라우저를 통해 토큰을 받아와 cli 통해 토큰 등록
func GetToken(ctx context.Context, config *oauth2.Config) (*oauth2.Token, error) {
	url := config.AuthCodeURL("state")
	fmt.Printf("Type the following url in to your broswer: %v\n", url)
	fmt.Println("Paste the token in the redirect URL and hit enter: ")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	return config.Exchange(ctx, code)
}
