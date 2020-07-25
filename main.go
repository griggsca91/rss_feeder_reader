package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"rss_feeder_reader/component"
	"rss_feeder_reader/customtheme"
	"rss_feeder_reader/model"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func getChannel(url string) (*model.Channel, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed model.RSS
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed.Channel, nil
}

func addFeedItemsToContainer(container *widget.Box, feed model.Channel) {
	container.Children = nil
	items := make([]fyne.CanvasObject, 0)
	for _, item := range feed.Items {
		feedItemRow := component.NewFeedItemRow(item)
		items = append(items, feedItemRow)
	}

	container.Children = items
}

func main() {
	app := app.New()

	feedContainer := widget.NewVBox()
	feedContainerScroller := widget.NewScrollContainer(feedContainer)
	t := customtheme.NewCustomTheme()
	app.Settings().SetTheme(t)

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
