package utils

import "fyne.io/fyne/v2"

type DiagonalLayout struct {
}

func (d *DiagonalLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pad := fyne.NewSize(30, 30)
	for i, obj := range objects {
		obj.Move(fyne.NewPos(float32(i)*pad.Width, float32(i)*pad.Height))
		obj.Resize(obj.MinSize())
	}
}

func (d *DiagonalLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	pad := fyne.NewSize(10, 10)
	w, h := float32(0), float32(0)

	for i, obj := range objects {
		min := obj.MinSize()

		w = float32(i)*pad.Width + min.Width
		h = float32(i)*pad.Height + min.Height
	}

	return fyne.NewSize(w, h)
}
