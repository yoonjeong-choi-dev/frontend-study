package feed

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Title     string `xml:"title"`
	URL       string `xml:"link"`
	Published string `xml:"pubDate"`
}

type RSS struct {
	Channel struct {
		Items []Item `xml:"item"`
	} `xml:"channel"`
	EntityTag string // Etag
}

func (r *RSS) Items() []Item {
	items := make([]Item, len(r.Channel.Items))
	copy(items, r.Channel.Items)
	return items
}

func (r *RSS) ParseURL(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if r.EntityTag != "" {
		req.Header.Add("ETag", r.EntityTag)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case http.StatusNotModified:
		// 컨텐츠가 변경되지 않았으므로, 클라이언트 캐싱 데이터 사용
	case http.StatusOK:
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		_ = res.Body.Close()

		r.Channel.Items = make([]Item, 0)
		err = xml.Unmarshal(b, r)
		if err != nil {
			return err
		}

		r.EntityTag = res.Header.Get("ETag")
	default:
		return fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}
	return nil
}
