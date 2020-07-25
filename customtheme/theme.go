package customtheme

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

var (
	purple            = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	orange            = &color.NRGBA{R: 198, G: 123, B: 0, A: 255}
	grey              = &color.Gray{Y: 123}
	ItemRowBackground = &color.RGBA{117, 117, 117, 255}
)

// CustomTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type CustomTheme struct {
}

func (CustomTheme) BackgroundColor() color.Color {
	return purple
}

func (CustomTheme) ButtonColor() color.Color {
	return color.Black
}

func (CustomTheme) DisabledButtonColor() color.Color {
	return color.White
}

func (CustomTheme) HyperlinkColor() color.Color {
	return orange
}

func (CustomTheme) TextColor() color.Color {
	return color.White
}

func (CustomTheme) DisabledTextColor() color.Color {
	return color.Black
}

func (CustomTheme) IconColor() color.Color {
	return color.White
}

func (CustomTheme) DisabledIconColor() color.Color {
	return color.Black
}

func (CustomTheme) PlaceHolderColor() color.Color {
	return grey
}

func (CustomTheme) PrimaryColor() color.Color {
	return orange
}

func (CustomTheme) HoverColor() color.Color {
	return orange
}

func (CustomTheme) FocusColor() color.Color {
	return orange
}

func (CustomTheme) ScrollBarColor() color.Color {
	return grey
}

func (CustomTheme) ShadowColor() color.Color {
	return &color.RGBA{0xcc, 0xcc, 0xcc, 0xcc}
}

func (CustomTheme) TextSize() int {
	return 12
}

func (CustomTheme) TextFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (CustomTheme) TextBoldFont() fyne.Resource {
	return theme.DefaultTextBoldFont()
}

func (CustomTheme) TextItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (CustomTheme) TextBoldItalicFont() fyne.Resource {
	return theme.DefaultTextBoldItalicFont()
}

func (CustomTheme) TextMonospaceFont() fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (CustomTheme) Padding() int {
	return 5
}

func (CustomTheme) IconInlineSize() int {
	return 20
}

func (CustomTheme) ScrollBarSize() int {
	return 10
}

func (CustomTheme) ScrollBarSmallSize() int {
	return 5
}

func NewCustomTheme() fyne.Theme {
	return &CustomTheme{}
}
