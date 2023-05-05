package feed

import (
	"context"
	"testing"
)

func TestRSS_ParseURL(t *testing.T) {
	rss := new(RSS)
	err := rss.ParseURL(context.Background(), "https://xkcd.com/rss.xml")
	if err != nil {
		t.Fatal(err)
	}

	if rss.entityTag == "" {
		t.Error("unexpected empty Etag")
	}

	etag := rss.entityTag
	t.Logf("RSS Etag: %s\n", etag)

	err = rss.ParseURL(context.Background(), "https://xkcd.com/rss.xml")
	if err != nil {
		t.Fatal(err)
	}

	if etag != rss.entityTag {
		t.Log("entity update -> new Etag on second request")
	}

	items := rss.Items()
	for _, item := range items {
		t.Logf("%#v\n", item)
	}
}
