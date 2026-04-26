package main

import (
	"06-handling-forms/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("forms.com")
	window := app.NewWindow("Forms Showcase")
	window.Resize(fyne.NewSize(500, 500))
	window.SetIcon(theme.HomeIcon())

	showLoginForm(window, app)
	window.ShowAndRun()
}

// login func
func showLoginForm(w fyne.Window, a fyne.App) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	errorLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Italic: true})
	errorLabel.Hide()

	loginBtn := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if username == "" || password == "" {
			errorLabel.SetText("Uername & password cannont be empty")
			errorLabel.Show()
		}

		if username == "username" && password == "password" {
			// show dashboard
			showDashboard(w, a)
		} else {
			dialog.ShowInformation("Login Error", "wrong Uername or password", w)
		}

	})

	// register button
	registerBtn := widget.NewButton("Register", func() {
		showRegisterForm(w, a)
	})

	content := container.NewVBox(
		usernameEntry,
		passwordEntry,
		errorLabel,
		loginBtn,
		registerBtn,
	)

	centeredForm := utils.NewFixedWidthCenter(content, 300)

	w.SetContent(centeredForm)
}

func showRegisterForm(w fyne.Window, a fyne.App) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	confirmPasswordEntry := widget.NewPasswordEntry()
	confirmPasswordEntry.SetPlaceHolder("Confirm Password")

	registerBtn := widget.NewButton("Register", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		confirmPassword := confirmPasswordEntry.Text

		if username == "" || password == "" || confirmPassword == "" {
			dialog.ShowInformation("Login Error", "inputs cannot be empty", w)
		} else {
			if password != confirmPassword {
				dialog.ShowInformation("Password", "password nd confirm password do not mach", w)
			} else {
				dialog.ShowInformation("Regiter", "account registered", w)
				showLoginForm(w, a)
			}
		}
	})

	loginBtn := widget.NewButton("Login", func() {
		showLoginForm(w, a)
	})

	content := container.NewVBox(
		usernameEntry,
		passwordEntry,
		confirmPasswordEntry,
		registerBtn,
		loginBtn,
	)

	centeredForm := utils.NewFixedWidthCenter(content, 300)

	w.SetContent(centeredForm)
}

func showDashboard(w fyne.Window, a fyne.App) {
	label := widget.NewLabelWithStyle(
		"Welcome to the Dashboard",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true})

	logoutBtn := widget.NewButton("LogOut", func() {
		showLoginForm(w, a)
	})

	content := container.NewVBox(
		label,
		logoutBtn,
	)

	w.SetContent(content)
}
