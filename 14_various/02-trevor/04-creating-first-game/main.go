package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Creating Your First Game with Ebitengine - Trevors-Tutorials.com #4")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetFullscreen(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(100)

	err := ebiten.RunGame(&game{})
	if err != nil {
		log.Fatal(err)
	}
}
