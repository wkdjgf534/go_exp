package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Getting Started with Ebitengine - Trevors-Tutorials.com #3")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
