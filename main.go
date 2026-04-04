package main

import (
	"os"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.guysan.goinvest", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	version := "0.0.1"
	window := gtk.NewApplicationWindow(app)
	window.SetTitle("GoInvest " + version)
	window.SetDefaultSize(1200, 800)

	header := gtk.NewCenterBox()

	clockLabel := gtk.NewLabel("")
	clockLabel.SetText("00:00:00")
	header.Append(clockLabel)

	rootBox := gtk.NewBox(gtk.OrientationVertical, 0)
	rootBox.Append(header)

	mainPanel := gtk.NewBox(gtk.OrientationHorizontal, 0)
	mainPanel.SetVExpand(true)
	rootBox.Append(mainPanel)

	// sidebar stuff
	sidebar := gtk.NewBox(gtk.OrientationVertical, 0)

	watchTitle := gtk.NewLabel("")
	watchTitle.SetMarkup("<span size='large' weight='bold'>Watchlist</span>")
	watchTitle.SetMarginTop(15)
	watchTitle.SetMarginBottom(15)
	sidebar.Append(watchTitle)

	listbox := gtk.NewListBox()
	listbox.SetSelectionMode(gtk.SelectionSingle)
	listbox.Append(gtk.NewLabel("AAPL  -  $35.20"))
	listbox.Append(gtk.NewLabel("MSFT  -  $29.15"))
	listbox.Append(gtk.NewLabel("GOOG  -  $256.10"))
	sidebar.Append(listbox)

	// search by ticker
	searchbox := gtk.NewBox(gtk.OrientationHorizontal, 5)
	searchbox.SetMarginTop(15)

	searchEntry := gtk.NewSearchEntry()
	searchEntry.SetPlaceholderText("Search ticker")
	searchEntry.SetHExpand(true)
	searchbox.Append(searchEntry)

	addButton := gtk.NewButtonWithLabel("Add")
	searchbox.Append(addButton)
	sidebar.Append(searchbox)

	mainPanel.Append(sidebar)

	contentPaned := gtk.NewPaned(gtk.OrientationVertical)
	contentPaned.SetPosition(500)

	chartArea := gtk.NewDrawingArea()
	chartArea.SetHExpand(true)
	chartArea.SetVExpand(true)

	chartArea.SetDrawFunc(func(da *gtk.DrawingArea, cr *cairo.Context, width, height int) {
		cr.SetSourceRGB(0.08, 0.09, 0.11)
		cr.Paint()

		cr.SetSourceRGB(0.1, 0.8, 0.4)
		cr.SetLineWidth(2.5)
		cr.MoveTo(0, float64(height)*0.8)
		cr.LineTo(float64(width)*0.5, float64(height)*0.4)
		cr.LineTo(float64(width), float64(height)*0.2)
		cr.Stroke()
	})
	contentPaned.SetStartChild(chartArea)

	newsBox := gtk.NewBox(gtk.OrientationVertical, 10)
	newsBox.SetMarginTop(10)
	newsBox.SetMarginStart(10)
	newsBox.SetMarginEnd(10)

	newsTitle := gtk.NewLabel("")
	newsTitle.SetMarkup("<span size='large' weight='bold'>Live Simulated News</span>")
	newsTitle.SetXAlign(0)
	newsBox.Append(newsTitle)

	news1 := gtk.NewLabel("09:30 AM - Market opens. Dow Jones down 20 points.")
	news1.SetXAlign(0)
	newsBox.Append(news1)

	news2 := gtk.NewLabel("09:45 AM - Tech sector shows early signs of volatility.")
	news2.SetXAlign(0)
	newsBox.Append(news2)

	newsScroll := gtk.NewScrolledWindow()
	newsScroll.SetChild(newsBox)
	newsScroll.SetMinContentHeight(150)
	contentPaned.SetEndChild(newsScroll)

	mainPanel.Append(contentPaned)

	window.SetChild(rootBox)
	window.Show()
}
