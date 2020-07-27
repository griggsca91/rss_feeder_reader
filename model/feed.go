package model

import "time"

type Feeder interface {
	GetFeedItems() []FeedItem
}

type FeedItem struct {
	Title       string
	Description string
	PubDate     time.Time
	Link        string
	Source      string
}
