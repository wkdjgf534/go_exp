package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Helllo, World!")

	w.SetContent(widget.NewLabel("Hello, world!"))
	w.ShowAndRun() // w.Show and a.Run()
}
