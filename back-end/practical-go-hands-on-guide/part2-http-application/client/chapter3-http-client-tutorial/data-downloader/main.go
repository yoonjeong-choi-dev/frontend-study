package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Must specify a HTTP URL to fetch")
		os.Exit(1)
	}

	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)
}
