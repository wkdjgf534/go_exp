package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("Keyboard.shortcuts.com")
	window := app.NewWindow("Keyboard Shortcuts")
	window.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("No shortcut is clicked")

	// check if the app is running in desktop mode

	if _, ok := app.(desktop.App); ok {
		// define shortcuts
		ctrlN := &desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: fyne.KeyModifierControl}     // ctrl + n
		ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl} // ctrl + tab
		ctrlP := &desktop.CustomShortcut{KeyName: fyne.KeyP, Modifier: fyne.KeyModifierControl}     // ctrl + p

		// ctrl + n
		window.Canvas().AddShortcut(ctrlN, func(shortcut fyne.Shortcut) {
			label.SetText("ctrl+n clicked")
			fmt.Println(shortcut)
		})

		// ctrl + tab
		window.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
			label.SetText("ctrl+tab clicked")
			fmt.Println(shortcut)
		})

		// ctrl + p
		window.Canvas().AddShortcut(ctrlP, func(shortcut fyne.Shortcut) {
			label.SetText("ctrl+p clicked")
			fmt.Println(shortcut)
		})
	} else {
		log.Println("app not running on desktop")
	}

	window.SetContent(label)
	window.ShowAndRun()
}
