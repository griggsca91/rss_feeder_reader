package model

import "encoding/xml"

type RSSv2 struct {
	xml.Name `xml:"rss"`
	Channel  Channel `xml:"channel"`
}
