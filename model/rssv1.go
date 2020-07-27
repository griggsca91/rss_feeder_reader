package model

import (
	"encoding/xml"
	"fmt"
	"time"
)

type RSSv1 struct {
	xml.Name `xml:"rss"`
	Channel  Channel `xml:"channel"`
}

type Channel struct {
	Title         string `xml:"title"`
	LastBuildDate string `xml:"lastBuildDate"`
	Items         []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
	Comments    string `xml:"comments"`
	Guid        string `xml:"guid"`
}

func (e Item) GetPubDate() (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", e.PubDate)
}

func (r RSSv1) GetFeedItems() []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range r.Channel.Items {
		fmt.Printf("rss FeedItem %+v \n", item)
		date, err := item.GetPubDate()
		if err != nil {
			date = time.Time{}
		}
		feedItem := FeedItem{
			Title:       item.Title,
			Description: item.Description,
			PubDate:     date,
			Link:        item.Link,
			Source:      r.Channel.Title,
		}
		feedItems = append(feedItems, feedItem)
	}

	return feedItems
}
