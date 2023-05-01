package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/pkg/errors"
)

// Client 서비스 클라이언트 인터페이스
type Client interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}) error
	Close() error
}

// firebaseClient Client 구현체
// Close() error 메서드는 임베딩한 firestore.Client 에 구현되어 있음
type firebaseClient struct {
	*firestore.Client
	collection string
}

func (f *firebaseClient) Get(ctx context.Context, key string) (interface{}, error) {
	data, err := f.Collection(f.collection).Doc(key).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}

	return data.Data(), err
}

func (f *firebaseClient) Set(ctx context.Context, key string, value interface{}) error {
	data := make(map[string]interface{})
	data[key] = value

	_, err := f.Collection(f.collection).Doc(key).Set(ctx, data)
	return errors.Wrap(err, "set failed")
}
