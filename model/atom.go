package model

import (
	"encoding/xml"
	"time"
)

type Atom struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string   `xml:"title"`
	ID      string   `xml:"id"`
	Link    []Link   `xml:"link"`
	Updated TimeStr  `xml:"updated"`
	Author  *Person  `xml:"author"`
	Entry   []*Entry `xml:"entry"`
}

type Entry struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Updated     string `xml:"updated"`
	Link        Link   `xml:"link"`
	Comments    string `xml:"comments"`
	Guid        string `xml:"guid"`
}

type Link struct {
	Rel      string `xml:"rel,attr,omitempty"`
	Href     string `xml:"href,attr"`
	Type     string `xml:"type,attr,omitempty"`
	HrefLang string `xml:"hreflang,attr,omitempty"`
	Title    string `xml:"title,attr,omitempty"`
	Length   uint   `xml:"length,attr,omitempty"`
}

type Person struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri,omitempty"`
	Email    string `xml:"email,omitempty"`
	InnerXML string `xml:",innerxml"`
}

type Text struct {
	Type string `xml:"type,attr"`
	Body string `xml:",chardata"`
}

type TimeStr string

func Time(t time.Time) TimeStr {
	return TimeStr(t.Format("2006-01-02T15:04:05-07:00"))
}

func (e Entry) GetPubDate() (t time.Time) {
	t, err := time.Parse("2006-01-02T15:04:05-07:00", e.PubDate)
	if err != nil {
		t, _ = time.Parse("2006-01-02T15:04:05-07:00", e.Updated)
	}

	return
}

func (a Atom) GetFeedItems() []FeedItem {
	feedItems := make([]FeedItem, 0)
	for _, item := range a.Entry {

		pubDate := item.GetPubDate()
		feedItem := FeedItem{
			Title:       item.Title,
			Description: item.Description,
			PubDate:     pubDate,
			Link:        item.Link.Href,
			Source:      a.Title,
		}
		feedItems = append(feedItems, feedItem)
	}

	return feedItems
}
