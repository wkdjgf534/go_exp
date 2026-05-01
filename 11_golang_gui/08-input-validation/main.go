package main

import (
	"08-input-validation/utils"
	"fmt"
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("validations.com")
	window := app.NewWindow("Validation Showcase")
	window.Resize(fyne.NewSize(500, 500))

	userNameEntry := widget.NewEntry()
	userNameEntry.SetPlaceHolder("username")
	userNameEntry.Validator = validation.NewRegexp("^[A-Za-z]{3,}$", "Username must be at least 3 characters")
	userNameError := widget.NewLabel("")
	userNameError.Hide()

	passwordEntry := widget.NewEntry()
	passwordEntry.SetPlaceHolder("password")
	passwordEntry.Validator = validation.NewRegexp("^[A-Za-z]{6,}$", "Password must be at least 6 characters")
	passwordError := widget.NewLabel("")
	passwordError.Hide()

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("pone number")
	phoneEntry.Validator = validation.NewRegexp(`^\+254\d{9}$`, "Invalid phone number")
	phoneError := widget.NewLabel("")
	phoneError.Hide()

	submitButton := widget.NewButton("submit", func() {
		username := userNameEntry.Text
		phone := phoneEntry.Text
		password := passwordEntry.Text

		dialog.ShowInformation(
			"Success",
			fmt.Sprintf("User registred:\n username: %s\n phone: %s\n password: %s", username, phone, password),
			window,
		)
	})

	submitButton.Disable()

	validateForm := func() {
		isValid := false

		if !regexp.MustCompile("^[A-Za-z]{3,}$").MatchString(userNameEntry.Text) {
			userNameError.SetText("username must contain at least 3 characters")
			userNameError.Show()
			isValid = false
		} else {
			isValid = true
			userNameError.Hide()
		}

		if !regexp.MustCompile("^[A-Za-z]{6,}$").MatchString(passwordEntry.Text) {
			passwordError.SetText("password must be at least 6 characters")
			passwordError.Show()
			isValid = false
		} else {
			isValid = true
			passwordError.Hide()
		}

		if !regexp.MustCompile(`^\+254\d{9}$`).MatchString(phoneEntry.Text) {
			phoneError.SetText("invalid phone number")
			phoneError.Show()
			isValid = false
		} else {
			isValid = true
			phoneError.Hide()
		}

		if isValid {
			submitButton.Enable()
		} else {
			submitButton.Disable()
		}
	}

	userNameEntry.OnChanged = func(s string) { validateForm() }
	passwordEntry.OnChanged = func(s string) { validateForm() }
	phoneEntry.OnChanged = func(s string) { validateForm() }

	form := container.NewVBox(
		userNameEntry,
		userNameError,
		phoneEntry,
		phoneError,
		passwordEntry,
		passwordError,
		submitButton,
	)

	centerForm := utils.NewFixedWidthCenter(form, 300)
	window.SetContent(centerForm)
	window.ShowAndRun()
}
