package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yoonjeong-choi-dev/moloco-study/back-end/go-network-programming/part4-service-architecture/chapter14-serverless/feed"
)

var (
	rssFeed feed.RSS
	feedURL = "https://xkcd.com/rss.xml"
)

// EventRequest unmarshalled request by AWS Lambda
type EventRequest struct {
	Previous bool `json:"previous"`
	All      bool `json:"all"`
}

type Item struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Published string `json:"published"`
}

// EventResponse response to be marshalled by AWS Lambda
type EventResponse struct {
	Items []Item `json:"items"`
}

// main Entry point of AWS Lambda
func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req EventRequest) (EventResponse, error) {
	// Default Response Data
	res := EventResponse{
		Items: []Item{{
			Title: "xkcd.com", URL: "https://xkcd.com",
		}},
	}

	// Fetch from xkcd rss
	if err := rssFeed.ParseURL(ctx, feedURL); err != nil {
		return res, err
	}

	// Only return the head data
	switch items := rssFeed.Items(); {
	case req.Previous && len(items) > 1:
		res.Items[0].Title = items[1].Title
		res.Items[0].URL = items[1].URL
		res.Items[0].Published = items[1].Published
	case len(items) > 0 && !req.All:
		res.Items[0].Title = items[0].Title
		res.Items[0].URL = items[0].URL
		res.Items[0].Published = items[0].Published
	case len(items) > 0 && req.All:
		res.Items = make([]Item, len(items))
		for i, item := range items {
			res.Items[i] = Item{
				Title:     item.Title,
				URL:       item.URL,
				Published: item.Published,
			}
		}
	}

	return res, nil
}
