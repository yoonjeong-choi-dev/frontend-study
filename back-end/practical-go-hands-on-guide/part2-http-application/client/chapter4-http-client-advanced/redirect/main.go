package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchRemoteResource(client *http.Client, url string) ([]byte, error) {
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func redirectPolicyFunc(req *http.Request, path []*http.Request) error {
	// path: [origin url, next url1, next url2,...]
	if len(path) >= 1 {
		return errors.New(fmt.Sprintf("\nStart from %s\nAttempted redirect to %s\n",
			path[0].URL, req.URL))
	}
	return nil
}

func main() {
	client := &http.Client{CheckRedirect: redirectPolicyFunc}
	body, err := fetchRemoteResource(client, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)
}
