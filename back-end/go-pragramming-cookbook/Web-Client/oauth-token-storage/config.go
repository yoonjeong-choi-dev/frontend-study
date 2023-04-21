package oauth_token_storage

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

type Config struct {
	*oauth2.Config
	Storage
}

func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	err = c.Storage.SetToken(token)
	return token, err
}

// TokenSource Oauth Client 생성 시, 필요한 TokenSource 객체를 생선하는 메서드
// => Custom Config 객체에 맞는 객체를 생성하도록 오버라이드
func (c *Config) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, token)
}

func (c *Config) GetToken(ctx context.Context) (*oauth2.Token, error) {
	token, err := c.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}

	url := c.AuthCodeURL("state")
	fmt.Printf("Type the following url in to your broswer: %v\n", url)
	fmt.Println("Paste the token in the redirect URL and hit enter: ")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	return c.Exchange(ctx, code)
}
