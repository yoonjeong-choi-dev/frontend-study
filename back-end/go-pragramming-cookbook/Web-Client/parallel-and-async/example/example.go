package main

import (
	"async"
	"fmt"
	"net/http"
)

func fetchAll(urls []string, c *async.Client) {
	for _, url := range urls {
		// 비동기 병렬 처리
		go c.AsyncGet(url)
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://golang.org",
		"https://www.github.com/yoonjeong-choi-dev",
	}

	client := async.CreateNewClient(http.DefaultClient, len(urls))
	fetchAll(urls, client)

	for i := 0; i < len(urls); i++ {
		// 모든 비동기 처리에 대해서 응답 대기 및 출력
		// 가장 먼저 응답 받는 요청에 대해서 출력
		select {
		case resp := <-client.Response:
			fmt.Printf("[%d] Status received from %s: %d\n", i, resp.Request.URL, resp.StatusCode)
		case err := <-client.Error:
			fmt.Printf("[%d] Error received: %s\n", err)
		}
	}
}
