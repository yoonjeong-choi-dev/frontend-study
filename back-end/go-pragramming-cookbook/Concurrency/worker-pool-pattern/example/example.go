package main

import (
	"context"
	"fmt"
	poolPattern "worker-pool-pattern"
)

func main() {
	cancel, in, out := poolPattern.CryptoWorkerPool(context.Background(), 10)
	defer cancel()

	numHashed := 10
	for i := 0; i < numHashed; i++ {
		in <- poolPattern.CryptoRequest{
			Op:   poolPattern.Hash,
			Text: []byte(fmt.Sprintf("secret message %d", i)),
		}
	}

	for i := 0; i < numHashed; i++ {
		res := <-out
		in <- poolPattern.CryptoRequest{
			Op:     poolPattern.Compare,
			Text:   res.Request.Text,
			Hashed: res.Result,
		}
	}

	for i := 0; i < numHashed; i++ {
		res := <-out
		if res.Err != nil {
			panic(res.Err)
		}

		fmt.Printf("Plain Text: '%s', matched: %v\n", string(res.Request.Text), res.Matched)
	}
}
