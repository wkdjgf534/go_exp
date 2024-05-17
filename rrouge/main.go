package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map GameMap
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	return g

}

func (g *Game) Update() error {
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the Map
	gd := NewGameData()
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x += 1 {
		for y := 0; y < gd.ScreenHeight; y += 1 {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
