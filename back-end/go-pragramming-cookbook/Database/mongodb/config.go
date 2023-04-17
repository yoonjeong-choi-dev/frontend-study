package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DB struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func Setup(ctx context.Context, host string) (*DB, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		cancel()
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		cancel()
		return nil, err
	}
	return &DB{client, ctx, cancel}, nil
}
