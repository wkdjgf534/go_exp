package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("fyne Collection")
	w := a.NewWindow("Fyne Collections Showcase")
	w.Resize(fyne.NewSize(600, 600))

	// list
	listData := []string{"Apple", "Banana", "Oranges", "Cherry", "Dates"}

	selectedItems := widget.NewLabel("No item selected")

	list := widget.NewList(
		func() int {
			return len(listData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(listData[i])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		selectedItems.SetText(listData[id])
	}

	// table
	userData := [][]string{
		{"Alice", "alice@example.com", "alice1234"},
		{"Ben", "ben@example.com", "ben1234"},
		{"Peter", "peter@example.com", "peter1234"},
	}

	table := widget.NewTable(
		func() (rows int, cols int) {
			return len(userData) + 1, 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row == 0 {
				headers := []string{"Username", "Email", "Password"}
				label.SetText(headers[id.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				label.SetText(userData[id.Row-1][id.Col])
			}
		},
	)

	table.SetColumnWidth(0, 120)
	table.SetColumnWidth(1, 200)
	table.SetColumnWidth(2, 150)

	// tree
	data := map[string][]string{
		"":           {"Fruits", "Vegetables"}, // root node
		"Fruits":     {"Apple", "Banana", "Orange"},
		"Vegetables": {"Carrot", "Brocolli", "Spinach"},
	}

	selectedNode := widget.NewLabel("No node is selected:")

	tree := widget.NewTree(
		func(id string) []string {
			return data[id]
		},
		func(id string) bool {
			return len(data[id]) > 0
		},
		func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id string, branch bool, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(id)
		},
	)

	tree.OnSelected = func(id string) {
		selectedNode.SetText(id)
	}
	// tabs holder
	tabs := container.NewAppTabs(
		container.NewTabItem("List", container.NewHBox(list, selectedItems)),
		container.NewTabItem("Table", table),
		container.NewTabItem("Tree", container.NewHBox(tree, selectedNode)),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}
