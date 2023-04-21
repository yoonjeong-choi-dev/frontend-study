package oauth_token_storage

import "golang.org/x/oauth2"

// Storage Generic Storage Interface for tokens
// => 다양한 저장소에 연결 가능
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(*oauth2.Token) error
}
