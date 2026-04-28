package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("icons.com")
	window := app.NewWindow("Icons Showcase")
	window.Resize(fyne.NewSize(500, 500))

	iconPath := "fyne-logo.png"

	iconData, err := os.ReadFile(iconPath)
	if err != nil {
		fyne.LogError("Failed to load the icon", err)
	}

	iconResource := fyne.NewStaticResource("icon", iconData)
	window.SetIcon(iconResource)

	home := widget.NewIcon(theme.HomeIcon())
	homeContainer := container.NewVBox(
		home,
		widget.NewLabelWithStyle("Home icon", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	mailSend := widget.NewIcon(theme.MailSendIcon())
	mailSendContainer := container.NewVBox(
		mailSend,
		widget.NewLabelWithStyle("Mail Send icon", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	account := widget.NewIcon(theme.AccountIcon())
	accountContainer := container.NewVBox(
		account,
		widget.NewLabelWithStyle("Account icon", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
	)

	// load images and display them
	image := canvas.NewImageFromFile(iconPath)
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(100, 100))

	// shapes
	rec := canvas.NewRectangle(color.White)
	rec.SetMinSize(fyne.NewSize(100, 100))

	// line
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 4

	homeGrid := container.NewGridWithColumns(3, homeContainer, mailSendContainer, accountContainer)

	window.SetContent(container.NewVBox(
		homeGrid,
		image,
		rec,
		line,
	))
	window.ShowAndRun()
}
