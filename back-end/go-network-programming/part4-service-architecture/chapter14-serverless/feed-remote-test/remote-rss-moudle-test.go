package main

import (
	"context"
	"fmt"
	"github.com/yoonjeong-choi-dev/moloco-study/back-end/go-network-programming/part4-service-architecture/chapter14-serverless/feed"
)

func main() {
	rss := new(feed.RSS)
	err := rss.ParseURL(context.Background(), "https://xkcd.com/rss.xml")
	if err != nil {
		panic(err)
	}

	fmt.Println("Response Items:")
	items := rss.Items()
	for _, item := range items {
		fmt.Printf("%#v\n", item)
	}
}
