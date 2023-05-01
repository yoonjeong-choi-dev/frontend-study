package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"time"
)

type Message struct {
	Message   string
	Timestamp time.Time
}

const MessageKey = "Message"

func (c *Controller) storeMessage(ctx context.Context, message string) error {
	m := &Message{
		Message:   message,
		Timestamp: time.Now(),
	}

	key := datastore.IncompleteKey(MessageKey, nil)
	_, err := c.store.Put(ctx, key, m)
	return err
}

func (c *Controller) getMessage(ctx context.Context, limit int) ([]*Message, error) {
	q := datastore.NewQuery(MessageKey).Order("-Timestamp").Limit(limit)

	ret := make([]*Message, 0)
	_, err := c.store.GetAll(ctx, q, &ret)
	return ret, err
}
