package entities

import "github.com/hajimehoshi/ebiten/v2"

// The base struct for all our moving, drawn entities
type Sprite struct {
	Img          *ebiten.Image
	X, Y, Dx, Dy float64
}
