package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	background color.RGBA
}

func newGame() *game {
	return &game{
		background: color.RGBA{0, 0, 255, 255},
	}
}

// Layout returns the size of the application. If the window is smaller or larger,
// the application is scaled while maintaining the correct aspect ratio.
func (g *game) Layout(outsideWith, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 270
}

// Update - updates the game state.
func (g *game) Update() error {
	return nil
}

// Draw - draws the game onto the screen.
func (g *game) Draw(screen *ebiten.Image) {
	// Draw background
	screen.Fill(g.background)
}
