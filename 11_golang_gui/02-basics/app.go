package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("fyne widgets")
	w := a.NewWindow("Fyne Widgets Showcace")
	w.Resize(fyne.NewSize(600, 600))

	// 1. label
	// method 1
	// title := widget.NewLabel("Fyne widgets")
	// title.Alignment = fyne.TextAlignCenter
	// title.TextStyle = fyne.TextStyle{Bold: true}

	// method 2
	title := widget.NewLabelWithStyle("Fyne widgets", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// --------

	// 2. buttons
	button := widget.NewButton("Update Title", func() {
		title.SetText("Button clicked")
	})

	button2 := widget.NewButtonWithIcon("Reset Title", theme.HomeIcon(), func() {
		title.SetText("Fyne widgets")
	})

	// buttons container
	buttonsContainer := container.New(layout.NewGridLayoutWithColumns(2), button, button2)

	// --------

	// 3. User Entry
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	// --------

	// 4. Checkbox
	checkbox := widget.NewCheck("I accept the terms", func(checked bool) {})

	// --------

	// 5. Radio button
	radioBtn := widget.NewRadioGroup([]string{"Man", "Female"}, func(selected string) {})

	// --------

	// 6. Text area
	textarea := widget.NewMultiLineEntry()
	textarea.SetPlaceHolder("Enter your message")

	// --------

	// 7. Slider
	slider := widget.NewSlider(0, 100)

	// --------

	// 8 Select
	sizeSelect := widget.NewSelect([]string{"Group A", "Group B", "Group C"}, func(s string) {})
	sizeSelect.SetSelected("Group A")

	// --------

	// 8 Select With Input
	sizeSelect2 := widget.NewSelectEntry([]string{"Group A", "Group B", "Group C"})

	// --------

	// 9 Progress bar
	progress := widget.NewProgressBar()

	go func() {
		for i := 0.0; i < 1.0; i += 0.1 {
			time.Sleep(time.Second)
			fyne.Do(func() {
				progress.SetValue(i)
			})
		}
	}()

	// 10 Toolbar
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("Document created")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			log.Println("Document cut")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			log.Println("Document copied")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			log.Println("Document pasted")
		}),
	)

	// 11 File upload
	fileUpload := widget.NewLabel("No file uploader")

	file := widget.NewButton("Open File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				fileUpload.SetText("Opened: " + reader.URI().String())
			}
		}, w)
	})

	// content holder
	content := container.New(
		layout.NewVBoxLayout(),
		toolbar,
		title,
		buttonsContainer,
		username,
		password,
		checkbox,
		radioBtn,
		radioBtn,
		textarea,
		slider,
		sizeSelect,
		sizeSelect2,
		progress,
		file,
		fileUpload,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
