package model

type Feeder interface {
  GetFeedItems() []FeedItem
}

type FeedItem struct {
  Title string
  Description string
  PubDate string
  Link string
  Source string
}
