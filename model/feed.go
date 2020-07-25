package model

type RSS struct {
	Channel Channel `xml:"channel"`
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
