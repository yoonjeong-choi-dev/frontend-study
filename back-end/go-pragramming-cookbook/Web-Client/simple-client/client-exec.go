package client

import (
	"fmt"
	"net/http"
)

func DoOperations(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	fmt.Printf("DoOperations Result Code from %s: %d\n", url, resp.StatusCode)
	return nil
}

func DefaultGetMethod(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	fmt.Printf("DefaultGetMethod Result Code from %s: %d\n", url, res.StatusCode)
	return nil
}
