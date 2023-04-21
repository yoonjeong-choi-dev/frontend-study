package oauth_token_storage

import (
	"context"
	"golang.org/x/oauth2"
)

// storageTokenSource oauth2.TokenSource 구현체
type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// Token implement the interface oauth2.TokenSource
func (s *storageTokenSource) Token() (*oauth2.Token, error) {
	// 스토리지에 저장된 토큰이 유효한 경우
	if token, err := s.Config.Storage.GetToken(); err == nil && token.Valid() {
		return token, err
	}

	// 유효한 토근이 없는 경우 새로 발급
	token, err := s.TokenSource.Token()
	if err != nil {
		return token, err
	}

	// 발급받은 토큰 저장
	err = s.Config.Storage.SetToken(token)
	return token, err
}

func StorageTokenSource(ctx context.Context, c *Config, token *oauth2.Token) oauth2.TokenSource {
	if token == nil || !token.Valid() {
		if saved, err := c.Storage.GetToken(); err == nil {
			token = saved
		}
	}

	source := c.Config.TokenSource(ctx, token)
	return &storageTokenSource{c, source}
}
