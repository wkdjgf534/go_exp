package goasteroids

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	baseMeteorVelocity  = 0.25                    // The base speed for meteors.
	meteorSpawnTime     = 100 * time.Millisecond  // How long before meteors spawn.
	meteorSpeedUpAmount = 0.1                     // How much do we speed a meteor up when it's timer runs out.
	meteorSpeedUpTime   = 1000 * time.Millisecond // How long to wait to speed up meteors.
)

// GameScene is the overall type for a game scene (e.g. TitleScene, GameScene, etc.).
type GameScene struct {
	player           *Player
	baseVelocity     float64         // The base velocity for items in the game.
	meteorCount      int             // The counter for meteors.
	meteorSpawnTimer *Timer          // The timer for spawning meteors.
	meteors          map[int]*Meteor // A map of meteors.
	meteorsForLevel  int             // # of meteors for a level.
	velocityTimer    *Timer          // The timer used for speeding up meteors.
}

// NewGameScene is a factory method for producing a new game. It's called once,
// when game play starts (and again when game play restarts).
func NewGameScene() *GameScene {
	g := &GameScene{
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
		baseVelocity: baseMeteorVelocity,
		velocityTimer: NewTimer(meteorSpeedUpTime),
		meteors: make(map[int]*Meteor),
		meteorCount: 0,
		meteorsForLevel: 2,
	}
	g.player = NewPlayer(g)

	return g
}

// Update updates all game scene elements for the next draw. It's called once per tick.
func (g *GameScene) Update(state *State) error {
	g.player.Update()

	g.spawnMeteors()

	for _, m := range g.meteors {
		m.Update()
	}

	g.speedUpMeteors()

	return nil
}

// Draw draws all game scene elements to the screen. It's called once per frame.
func (g *GameScene) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	// Draw meteors.
	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

// Layout is necessary to satisfy interface requirements from ebiten.
func (g *GameScene) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return outsideWidth, outsideHeight
}

// spawnMeteors creates meteors, up to the maximum for a level.
func (g *GameScene) spawnMeteors() {
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		if len(g.meteors) < g.meteorsForLevel && g.meteorCount < g.meteorsForLevel {
			m := NewMeteor(g.baseVelocity, g, len(g.meteors)-1)
			g.meteorCount++
			g.meteors[g.meteorCount] = m
		}
	}
}

// speedUpMeteors makes meteors move faster over time.
func (g *GameScene) speedUpMeteors() {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady(){
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}
}
