package main

import (
	_ "image/png"
	"log"

	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map         GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, tags := InitializeWorld(g.Map.CurrentLevel)
	g.WorldTags = tags
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0
	return g

}

func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 20 {
		TryMovePlayer(g)
	}

	if g.Turn == MonsterTurn {
		UpdateMonster(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
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
