package model

import (
	"encoding/xml"
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
	return time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", e.PubDate)
}

func (r RSSv1) GetFeedItems() []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range r.Channel.Items {
		date, _ := item.GetPubDate()
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
