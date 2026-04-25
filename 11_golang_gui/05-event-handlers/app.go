package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("event.handlers.com")
	window := app.NewWindow("Event Handlers Showcase")
	window.Resize(fyne.NewSize(500, 500))
	window.SetIcon(theme.HomeIcon())

	// initialize labels
	entryLabel := widget.NewLabel("Entry")
	passwordLabel := widget.NewLabel("Password")
	checkedLabel := widget.NewLabel("Checkbox: Not checked")
	radioLabel := widget.NewLabel("Radio: None selected")
	sliderLabel := widget.NewLabel("Slider: ")
	selectLabel := widget.NewLabel("Selected: ")

	// entry
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Full name")
	entry.OnChanged = func(s string) {
		entryLabel.SetText(s)
	}

	// checkbox
	checkbox := widget.NewCheck("Accept Trems", func(checked bool) {
		if checked {
			checkedLabel.SetText("Checkbox: checked")
			return
		}

		checkedLabel.SetText("Checkbox: Not checked")
	})

	// password
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Enter password")
	password.OnChanged = func(s string) {
		passwordLabel.SetText(s)
	}

	// radio button
	radio := widget.NewRadioGroup([]string{"Male", "Female"}, func(s string) {
		radioLabel.SetText(s)
	})

	// slider
	slider := widget.NewSlider(0, 100)
	slider.OnChanged = func(f float64) {
		sliderLabel.SetText(fmt.Sprintf("Slider: %.2f", f))
	}

	// select
	selectDrop := widget.NewSelect([]string{"maths", "english", "serbian"}, func(s string) {
		selectLabel.SetText(s)
	})

	// submit button
	submitBtn := widget.NewButton("Submit", func() {
		fmt.Println("Form submited")
		fmt.Println("Entry: ", entry.Text)
		fmt.Println("Password: ", passwordLabel.Text)
	})

	// form
	form := container.NewVBox(
		widget.NewLabel("Form"),
		entry,
		password,
		checkbox,
		radio,
		slider,
		selectDrop,
		submitBtn,
	)

	// responses
	res := container.NewVBox(
		widget.NewLabel("Form"),
		entryLabel,
		passwordLabel,
		checkedLabel,
		radioLabel,
		sliderLabel,
		selectLabel,
	)

	// content
	content := container.NewHSplit(
		form,
		res,
	)

	window.SetContent(content)
	window.ShowAndRun()
}
