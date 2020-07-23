package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

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

func getChannel(url string) (*Channel, error) {
	resp, err := http.Get(url)
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

func addFeedItemsToContainer(container *widget.Box, feed Channel) {
	container.Children = nil
	items := make([]fyne.CanvasObject, 0)
	for _, item := range feed.Items {
		hBox := widget.NewHBox(
			widget.NewLabel(item.Title),
		)
		items = append(items, hBox)
	}

	container.Children = items
}

func main() {
	app := app.New()

	feedContainer := widget.NewVBox()
	feedContainerScroller := widget.NewScrollContainer(feedContainer)

	w := app.NewWindow("Hello")

	refreshButton := widget.NewButton("Refresh", func() {
		feed, err := getChannel("https://hnrss.org/newest")
		log.Println("got the feed")
		if err != nil {
			log.Fatalf("Error getting feed %v", err)
		}

		addFeedItemsToContainer(feedContainer, *feed)

		feedContainer.Refresh()
		feedContainerScroller.Refresh()
	})
	rootContainer := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(refreshButton, nil, nil, nil),
		refreshButton,
		feedContainerScroller,
	)

	w.SetContent(rootContainer)

	w.ShowAndRun()
}
