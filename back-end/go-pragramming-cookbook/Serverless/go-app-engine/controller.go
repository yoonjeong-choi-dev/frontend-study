package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
	store *datastore.Client
}

func (c *Controller) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method(only GET supported)", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form data", http.StatusBadRequest)
		return
	}

	if message := r.FormValue("message"); message != "" {
		if err := c.storeMessage(ctx, message); err != nil {
			log.Printf("failed to store message: %v\n", err)
			http.Error(w, "failed to store message", http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintln(w, "Saved Messages:")
	messages, err := c.getMessage(ctx, 10)
	if err != nil {
		log.Printf("failed to get messages: %v\n", err)
		http.Error(w, "failed to get messages", http.StatusInternalServerError)
		return
	}

	for _, message := range messages {
		fmt.Fprintln(w, message.Message)
	}
}
