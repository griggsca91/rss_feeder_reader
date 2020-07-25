package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"rss_feeder_reader/component"
	"rss_feeder_reader/customtheme"
	"rss_feeder_reader/model"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func getFeed(url string) (*model.Channel, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header["user-agent"] = []string{"not a bot"}
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	contentType := resp.Header.Get("content-type")
	fmt.Println("content-type", contentType)
	if strings.Contains(contentType, "application/atom+xml") {
		fmt.Println("is Atom")
	} else if strings.Contains(contentType, "application/rss+xml") {
		fmt.Println("is rss")
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed model.RSSv1
	var atomFeed model.Atom
	err = xml.Unmarshal(body, &atomFeed)
	if err != nil {
		return nil, err
	}
	fmt.Println(atomFeed)

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

func createMenuItems() *fyne.MainMenu {
	setURLMenuItem := fyne.NewMenuItem("Set Menu Item", func() {
		log.Println("Set URL Menu Item")
	})

	return fyne.NewMainMenu(
		fyne.NewMenu("Edit", setURLMenuItem),
	)
}

func main() {
	app := app.NewWithID("rss_feeder_reader")
	app.SetIcon(theme.FyneLogo())

	feedContainer := widget.NewVBox()
	feedContainerScroller := widget.NewScrollContainer(feedContainer)
	t := customtheme.NewCustomTheme()
	app.Settings().SetTheme(t)

	w := app.NewWindow("Feeder Reader")
	setURLMenuItem := fyne.NewMenuItem("Set Menu Item", func() {
		log.Println("Set URL Menu Item")
	})

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Edit", setURLMenuItem),
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()

	refreshButton := widget.NewButton("Refresh", func() {
		feed, err := getFeed("https://reddit.com/r/programming/.rss")
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
	app.Preferences().SetString("URL", "https://raw.githubusercontent.com/griggsca91/rss_feeder_reader_list/master/sources.txt")
}
