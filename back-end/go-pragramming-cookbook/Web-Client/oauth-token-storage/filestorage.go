package oauth_token_storage

import (
	"encoding/json"
	"errors"
	"golang.org/x/oauth2"
	"log"
	"os"
	"sync"
)

// FileStorage 파일시스템 저장소를 사용하는 Storage 구현체
type FileStorage struct {
	Path  string
	mutex sync.RWMutex
}

func (f *FileStorage) GetToken() (*oauth2.Token, error) {
	f.mutex.RLock()
	defer f.mutex.RUnlock()

	file, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var token *oauth2.Token
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&token)
	if err != nil {
		return nil, err
	}

	log.Println("Get from storage")
	return token, nil
}

func (f *FileStorage) SetToken(token *oauth2.Token) error {
	if token == nil || !token.Valid() {
		return errors.New("bad token")
	}

	f.mutex.Lock()
	defer f.mutex.Unlock()

	file, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err == nil {
		log.Println("Saved to storage")
	}

	return err
}
