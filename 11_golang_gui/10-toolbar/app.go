package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("Toolbar.com")
	window := app.NewWindow("Toolbar Showcase")
	window.Resize(fyne.NewSize(500, 200))

	// label
	statusLabel := widget.NewLabel("Not clicked")

	// toolbar
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			statusLabel.SetText("document was clicked")
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			statusLabel.SetText("document was deleted")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.CancelIcon(), func() {
			statusLabel.SetText("document was canceled")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			statusLabel.SetText("document was saved")
		}),
	)

	window.SetContent(container.NewBorder(toolbar, statusLabel, nil, nil))
	window.ShowAndRun()
}
