package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("App.preferences.com")
	window := app.NewWindow("App preferences Showcase")
	window.Resize(fyne.NewSize(400, 400))

	prefs := app.Preferences()

	// label
	label := widget.NewLabel("")
	label.SetText("No mode selected")

	// username
	username := widget.NewEntry()
	username.SetPlaceHolder("username")
	username.SetText(prefs.StringWithFallback("username", ""))

	// darkmode
	darkModeCheck := widget.NewCheck("Dark Mode", nil)
	darkModeCheck.SetChecked(prefs.BoolWithFallback("darkmode", false))
	if darkModeCheck.Checked {
		label.SetText("Dark Mode")
	} else {
		label.SetText("Light Mode")
	}

	username.OnChanged = func(s string) {
		prefs.SetString("username", s)
	}

	darkModeCheck.OnChanged = func(b bool) {
		prefs.SetBool("darkmode", b)
		if b {
			label.SetText("Dark Mode")
		} else {
			label.SetText("Light Mode")
		}
	}

	content := container.NewVBox(
		label,
		username,
		darkModeCheck,
	)

	window.SetContent(content)
	window.ShowAndRun()
}
