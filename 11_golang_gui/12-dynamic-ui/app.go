package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.NewWithID("Data binding")
	window := app.NewWindow("Data binding Showcase")
	window.Resize(fyne.NewSize(500, 400))

	// two way string data binding
	str := binding.NewString()
	str.Set("Hello Fyne")

	label := widget.NewLabelWithData(str)
	entry := widget.NewEntryWithData(str)

	// float slider
	floatVal := binding.NewFloat()
	floatVal.Set(40.0)

	formattedStr := binding.FloatToStringWithFormat(floatVal, "Value: %.2f")
	floatLabel := widget.NewLabelWithData(formattedStr)

	slider := widget.NewSlider(0, 100)
	slider.Bind(floatVal)

	// list
	listData := binding.NewStringList()
	listData.Set([]string{"Item 1", "Item 2", "Item 3"})

	list := widget.NewListWithData(listData,
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			strItem, ok := di.(binding.String)
			if !ok {
				log.Println("Item is not a binding string")
			}
			co.(*widget.Label).Bind(strItem)
		},
	)

	newItem := widget.NewEntry()
	addButton := widget.NewButton("Add Item", func() {
		itemText := newItem.Text
		if itemText != "" {
			listData.Append(itemText)
			newItem.SetText("")
		}
	})

	labelContainer := container.NewGridWithColumns(2, label, entry)
	floatContainer := container.NewGridWithColumns(2, floatLabel, slider)
	listContainer := container.NewGridWithColumns(2, newItem, addButton)

	layout := container.NewVSplit(container.NewVBox(labelContainer, floatContainer), list)

	content := container.NewBorder(listContainer, nil, nil, nil, layout)

	window.SetContent(content)
	window.ShowAndRun()
}
