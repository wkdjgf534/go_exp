package main

import (
	"16-custom-widgets/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("custom.widget.com")
	window := app.NewWindow("Custom Widget")
	window.Resize(fyne.NewSize(250, 200))

	counter := utils.NewCustomCounter()
	content := container.NewVBox(
		widget.NewLabel("Custom counter widget"),
		counter,
	)

	window.SetContent(content)
	window.ShowAndRun()
}
