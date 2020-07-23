package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	LastBuildDate string `xml:"lastBuildDate"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	PubDate string `xml:"pubDate"`
	Link string `xml:"link"`
	Comments string `xml:"comments"`
	Guid string `xml:"guid"`
}

func getChannel(url string)  (*Channel, error) {
	resp, err := http.Get("https://hnrss.org/frontpage")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSS
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed.Channel, nil
}

func main() {
	app := app.New()

	channel, err := getChannel("https://hnrss.org/frontpage")
	if err != nil {
		log.Fatalf("Error getting feed %v", err)
	}


	w := app.NewWindow("Hello")
	rootContainer := widget.NewVBox()
	for _, item := range channel.Items {
		hBox := widget.NewHBox(
			widget.NewLabel(item.Title),
		)
		rootContainer.Append(hBox)
	}


	w.SetContent(rootContainer)

	w.ShowAndRun()
}
