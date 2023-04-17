package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// MongoStorage ItemStorage 구현체
type MongoStorage struct {
	*mongo.Client
	DB         string
	Collection string
	Cancel     context.CancelFunc
}

func NewMongoStorage(ctx context.Context, host, db, collection string) (*MongoStorage, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}

	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	return &MongoStorage{client, db, collection, cancel}, nil
}

func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	collection := m.Client.Database(m.DB).Collection(m.Collection)

	var item Item
	if err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (m *MongoStorage) Save(ctx context.Context, item *Item) error {
	collection := m.Client.Database(m.DB).Collection(m.Collection)
	_, err := collection.InsertOne(ctx, item)
	return err
}
