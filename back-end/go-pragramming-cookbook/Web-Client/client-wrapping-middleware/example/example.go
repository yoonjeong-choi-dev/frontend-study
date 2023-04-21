package main

import (
	"fmt"
	"wrapping"
)

func main() {
	urls := []string{"https://www.google.com", "https://github.com"}

	client := wrapping.Setup()

	for _, url := range urls {
		res, err := client.Get(url)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Response code from %s: %d\n", url, res.StatusCode)
	}
}
