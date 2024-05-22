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

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		if !tile.Blocked {
			pos.X += x
			pos.Y += y
			level.PlayerVisible.Compute(level, pos.X, pos.Y, 8)
		}
	}
	if x != 0 || y != 0 {
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
