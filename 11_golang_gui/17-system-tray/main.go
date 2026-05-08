package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("sysytray")
	w := a.NewWindow("SysTray")

	// load local image
	iconPath := "fyne-logo.png"
	iconData, err := os.ReadFile(filepath.Clean(iconPath))
	if err != nil {
		panic(err)
	}

	localIcon := fyne.NewStaticResource("trayIcon", iconData)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp", fyne.NewMenuItem("show", func() {
			w.Show()
		}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(localIcon)
	}

	w.SetContent(widget.NewLabel("Fyne System Tray"))

	w.SetCloseIntercept(func() {
		w.Hide()
	})

	w.ShowAndRun()
}
