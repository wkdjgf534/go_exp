package main

import (
	"15-custom-layouts/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("custom.layout.com")
	window := app.NewWindow("Custom Layout")
	window.Resize(fyne.NewSize(250, 200))

	labels := []fyne.CanvasObject{
		widget.NewLabel("item 1"),
		widget.NewLabel("item 2"),
		widget.NewLabel("item 3"),
		widget.NewLabel("item 4"),
		widget.NewLabel("item 5"),
	}
	content := container.New(&utils.DiagonalLayout{}, labels...)

	window.SetContent(content)
	window.ShowAndRun()
}
