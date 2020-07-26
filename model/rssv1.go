package model

import (
	"encoding/xml"
	"fmt"
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

func (r RSSv1) GetFeedItems() []FeedItem {
  feedItems := make([]FeedItem, 0)
  for _, item := range r.Channel.Items {
    fmt.Printf("rss FeedItem %+v \n", item)
    feedItem := FeedItem {
      Title: item.Title,
      Description: item.Description,
      PubDate: item.PubDate,
      Link: item.Link,
      Source: r.Channel.Title,
    }
    feedItems = append(feedItems, feedItem)
  }

  return feedItems
}
