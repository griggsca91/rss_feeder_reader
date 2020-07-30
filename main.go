package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"rss_feeder_reader/component"
	"rss_feeder_reader/customtheme"
	"rss_feeder_reader/model"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func getFeeds(url string) ([]string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(body), "\n"), nil
}

func getFeed(url string) (model.Feeder, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header["user-agent"] = []string{"not a bot"}
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	if strings.Contains(contentType, "application/atom+xml") {
		var atomFeed model.Atom
		if err = xml.Unmarshal(body, &atomFeed); err != nil {
			return nil, err
		}
		return &atomFeed, nil

	} else if strings.Contains(contentType, "application/rss+xml") {
		var rssFeed model.RSSv1
		if err = xml.Unmarshal(body, &rssFeed); err != nil {
			return nil, err
		}

		return &rssFeed, nil
	}

	return nil, fmt.Errorf("Invalid feed type: %s", contentType)
}

type ItemSlice []model.FeedItem

func (p ItemSlice) Len() int           { return len(p) }
func (p ItemSlice) Less(i, j int) bool { return p[i].PubDate.Before(p[j].PubDate) }
func (p ItemSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort is a convenience method.
func (p ItemSlice) Sort() { sort.Sort(p) }

var feedItems ItemSlice

func addFeedItemsToContainer(container *widget.Box, feed model.Feeder) {
	for _, item := range feed.GetFeedItems() {
		feedItems = append(feedItems, item)
	}

	feedItems.Sort()

	container.Children = nil
	for _, item := range feedItems {
		feedItemRow := component.NewFeedItemRow(item)
		container.Children = append(container.Children, feedItemRow)
	}
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

	refreshButton := widget.NewButton("Refresh", func() {
		url := app.Preferences().String("URL")
		feeds, err := getFeeds(url)
		if err != nil {
			log.Fatalf("Error getting list of feeds %v", err)
		}

		start := time.Now()
		var wg sync.WaitGroup
		for _, feedURL := range feeds {
			if feedURL == "" {
				continue
			}
			feed, err := getFeed(feedURL)
			if err != nil {
				log.Printf("Error getting feed %v", err)
				return
			}

			addFeedItemsToContainer(feedContainer, feed)
		}
		wg.Wait()

		fmt.Println("finished getting all the feeds", time.Since(start))
		feedContainer.Refresh()
		feedContainerScroller.Refresh()
	})
	rootContainer := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(refreshButton, nil, nil, nil),
		refreshButton,
		feedContainerScroller,
	)

	w.SetContent(rootContainer)

	app.Preferences().SetString("URL", "https://raw.githubusercontent.com/griggsca91/rss_feeder_reader_list/master/sources.txt")
	w.ShowAndRun()
}
