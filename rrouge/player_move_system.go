package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	X = 0
	Y = 0
)

func TryMovePlayer(g *Game) {
	players := g.WorldTags["players"]
	x := X
	y := Y
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		pos.X += x
		pos.Y += y
	}
}