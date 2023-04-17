package storage

import (
	"context"
	"fmt"
)

func storageExecExample(s ItemStorage) error {
	ctx := context.Background()

	items := []Item{
		{
			Name:   "MacBook",
			Price:  4000,
			OnSale: false,
		},
		{
			Name:   "IPhone",
			Price:  2000,
			OnSale: true,
		},
	}

	for _, i := range items {
		if err := s.Save(ctx, &i); err != nil {
			return err
		}
	}

	for _, i := range items {
		item, err := s.GetByName(ctx, i.Name)
		if err != nil {
			return err
		}
		fmt.Printf("Result with name 'name': %#v\n", item)
	}

	return nil
}

func MongoExecExample() error {
	ctx := context.Background()
	m, err := NewMongoStorage(ctx, "mongodb://localhost", "gocookbook", "fancyStore")
	if err != nil {
		return err
	}
	defer m.Cancel()

	if err := storageExecExample(m); err != nil {
		return err
	}

	// Clean up
	if err := m.Client.Database(m.DB).Collection(m.Collection).Drop(ctx); err != nil {
		return err
	}

	return nil
}
