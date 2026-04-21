package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("layouts.com")
	window := app.NewWindow("Layouts Showcase")
	window.Resize(fyne.NewSize(500, 500))
	window.SetIcon(theme.HomeIcon())

	// vertical box
	vbox := container.NewVBox(
		widget.NewLabel("One"),
		widget.NewLabel("Two"),
		widget.NewLabel("Three"),
		widget.NewLabel("Four"),
	)

	// vertical box
	// hbox := container.New(layout.NewHBoxLayout())
	hbox := container.NewHBox(
		widget.NewLabel("One"),
		widget.NewLabel("Two"),
		widget.NewLabel("Three"),
		widget.NewLabel("Four"),
	)

	// grid
	grid := container.NewGridWithRows(2,
		widget.NewLabel("One"),
		widget.NewLabel("Two"),
		widget.NewLabel("Three"),
		widget.NewLabel("Four"),
	)

	// border layout
	border := container.NewBorder(
		widget.NewLabel("Top"),
		widget.NewLabel("Bottom"),
		widget.NewLabel("Left"),
		widget.NewLabel("Right"),
	)

	// form layout
	username := widget.NewLabel("username")
	usernameEntry := widget.NewEntry()
	password := widget.NewLabel("password")
	passwordEntry := widget.NewEntry()

	form := container.New(layout.NewFormLayout(), username, usernameEntry, password, passwordEntry)

	// center layout
	centerLayout := container.NewCenter(
		widget.NewLabel("My content center layout"),
	)

	// stack layout (max)
	stackLayout := container.NewStack(
		widget.NewLabel("My content stack layout"),
	)

	content := container.New(
		layout.NewVBoxLayout(),
		vbox,
		hbox,
		grid,
		border,
		form,
		centerLayout,
		stackLayout,
	)

	window.SetContent(content)
	window.ShowAndRun()
}
