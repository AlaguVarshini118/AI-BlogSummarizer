package utils

import (
	"github.com/mmcdole/gofeed"
)

func FetchArticles(feedURL string) ([]*gofeed.Item, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		return nil, err
	}
	return feed.Items, nil
}
