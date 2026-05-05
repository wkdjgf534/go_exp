package utils

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CustomCounter struct {
	widget.BaseWidget
	value int
	label *widget.Label
	add   *widget.Button
	sub   *widget.Button
}

func NewCustomCounter() *CustomCounter {
	c := &CustomCounter{}
	c.label = widget.NewLabel("0")
	c.add = widget.NewButton("+", func() {
		c.value++
		c.label.SetText(fmt.Sprintf("%d", c.value))
	})
	c.sub = widget.NewButton("-", func() {
		c.value--
		c.label.SetText(fmt.Sprintf("%d", c.value))
	})

	c.ExtendBaseWidget(c)
	return c
}

// - 0 +
func (c *CustomCounter) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewHBox(c.sub, c.label, c.add))
}
