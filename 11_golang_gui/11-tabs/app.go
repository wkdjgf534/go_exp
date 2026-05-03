package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("Tabs.com")
	window := app.NewWindow("Tabs Showcase")
	window.Resize(fyne.NewSize(500, 200))

	// inbox
	data := []string{"Email 1", "Email 2", "Email 3", "Email 4"}

	inbox := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[lii])
		},
	)

	email := widget.NewEntry()
	email.SetPlaceHolder("email")

	subject := widget.NewMultiLineEntry()
	subject.SetPlaceHolder("message")

	// sent emails
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: email},
			{Text: "Message", Widget: subject},
		},
		OnSubmit: func() {
			log.Println("Form submited")
		},
	}

	// tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Inbox", theme.MailComposeIcon(), inbox),
		container.NewTabItemWithIcon("Sent", theme.MailComposeIcon(), form),
	)

	window.SetContent(tabs)
	window.ShowAndRun()
}
