package component

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"rss_feeder_reader/customtheme"
	"rss_feeder_reader/model"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func (f *FeedItemRow) CreateRenderer() fyne.WidgetRenderer {
	f.ExtendBaseWidget(f)
	lay := layout.NewHBoxLayout()

	return &feedItemRowRenderer{
		feedItemRow: f,
		layout:      lay,
		objects: []fyne.CanvasObject{
			widget.NewLabel(f.FeedItem.Source),
			widget.NewLabel(f.FeedItem.PubDate.Local().String()),
			widget.NewLabel(f.FeedItem.Title),
		},
	}
}

type FeedItemRow struct {
	*widget.BaseWidget
	FeedItem      model.FeedItem
	tapped        bool
	hovered       bool
	background    color.Color
	mouseDownTime time.Time
	Selected      bool
}

func (w FeedItemRow) Hide() {
	w.BaseWidget.Hide()
}

func (f FeedItemRow) Tapped(_ *fyne.PointEvent) {
	//log.Printf("I've been tapped title: %s link: %s \n", f.Title, f.Link)
	parsedUrl, err := url.Parse(f.FeedItem.Link)
	if err != nil {
		log.Println(err)
		return
	}
	fyne.CurrentApp().OpenURL(parsedUrl)
}

// This essentially identifies a "click session" so we can solve the problem of when a
// user clicks on a row, moves the mouse out, and unclicks out of the row, mouse down out
// of the row, and then moves back in.  We know it's not the same click at that point and
// we don't display the hover+click background color
var (
	globalMouseDownTime time.Time
)

func (f *FeedItemRow) MouseIn(m *desktop.MouseEvent) {
	//log.Println("Mouse In", m.Button, desktop.LeftMouseButton, m, f.Title)
	//log.Println("globalMouseDownTime", globalMouseDownTime, "f.mouseDownTime", f.mouseDownTime)
	f.hovered = true
	if f.tapped && !f.mouseDownTime.Equal(globalMouseDownTime) {
		f.tapped = false
	}
	f.Refresh()
}
func (f *FeedItemRow) MouseOut() {
	f.hovered = false
	f.Refresh()
}

func (f *FeedItemRow) MouseMoved(m *desktop.MouseEvent) {}

func (f *FeedItemRow) MouseUp(m *desktop.MouseEvent) {
	if f.tapped {
		f.tapped = false
		f.Refresh()
	}
}

func (f *FeedItemRow) MouseDown(m *desktop.MouseEvent) {
	if m.Button == desktop.LeftMouseButton {
		globalMouseDownTime = time.Now()
		f.mouseDownTime = globalMouseDownTime
		f.tapped = true
		f.Refresh()
	}
}

func NewFeedItemRow(item model.FeedItem) *FeedItemRow {
	return &FeedItemRow{
		BaseWidget: &widget.BaseWidget{},
		FeedItem:   item,
		background: customtheme.ItemRowBackground,
	}
}

// Refresh updates this box to match the current theme
func (f FeedItemRow) Refresh() {

	if f.Selected {
		f.background = color.RGBA{250, 0, 0, 1}
		log.Println("FeedItemRow selected", f.FeedItem.Title, f.background)
	} else {
		f.background = customtheme.ItemRowBackground
	}

	f.BaseWidget.Refresh()
}

type feedItemRowRenderer struct {
	layout      fyne.Layout
	feedItemRow *FeedItemRow

	objects []fyne.CanvasObject
}

func (b *feedItemRowRenderer) MinSize() fyne.Size {
	return b.layout.MinSize(b.Objects())
}

func (b *feedItemRowRenderer) Layout(size fyne.Size) {
	b.layout.Layout(b.Objects(), size)
}

func (b *feedItemRowRenderer) BackgroundColor() color.Color {
	if b.feedItemRow.Selected {
		return color.RGBA{250, 0, 0, 1}
	} else {
		return customtheme.ItemRowBackground
	}
}

func (b *feedItemRowRenderer) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *feedItemRowRenderer) Destroy() {}

func (b *feedItemRowRenderer) Refresh() {
	for _, child := range b.Objects() {
		child.Refresh()
	}

	if b.feedItemRow.Selected {
		fmt.Println("Refreshing feeditemrow", b.feedItemRow.FeedItem.Title)
	}
	canvas.Refresh(b.feedItemRow)
}
