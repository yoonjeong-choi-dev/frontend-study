package gcp

import (
	"encoding/json"
	"github.com/yoonjeong-choi-dev/moloco-study/back-end/go-network-programming/part4-service-architecture/chapter14-serverless/feed"
	"log"
	"net/http"
)

var (
	rssFeed feed.RSS
	feedURL = "https://xkcd.com/rss.xml"
)

type EventRequest struct {
	Previous bool `json:"previous"`
	All      bool `json:"all"`
}

type Item struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Published string `json:"published"`
}

type EventResponse struct {
	Items []Item `json:"items"`
}

// MainHandler needs encoding/decoding process for GCP
// cf. AWS Lambda supports encoding/decoding of response/request
func MainHandler(w http.ResponseWriter, r *http.Request) {
	var req EventRequest

	// Default Response Data
	res := EventResponse{
		Items: []Item{{
			Title: "xkcd.com", URL: "https://xkcd.com",
		}},
	}

	defer func() {
		w.Header().Set("Content-Type", "application/json")
		out, _ := json.Marshal(&res)
		_, _ = w.Write(out)
	}()

	// Unmarshal request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("decoding request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rssFeed.ParseURL(r.Context(), feedURL); err != nil {
		log.Printf("parsing feed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
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
}
