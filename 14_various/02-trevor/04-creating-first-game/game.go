package main

import (
	"embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed asset
var assetFS embed.FS

const (
	screenW, screenH = 270, 152
	playerW, playerH = 16, 16
)

func loadImage(assetPath string) *ebiten.Image {
	f, err := assetFS.Open(assetPath)
	if err != nil {
		log.Panic(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

type game struct {
	// Background color.
	backgroundColor color.RGBA

	// Player image.
	playerImg *ebiten.Image

	// Player position.
	playerX, playerY float64

	// Player velocity.
	playerVelocity float64

	// Game attributes.
	fullscreen  bool
	initialized bool
}

// reset game state.
func (g *game) reset() {
	// Center player on screen.
	g.playerX = screenW/2 - playerW/2
	g.playerY = screenH/2 - playerH/2
}

// initialize sets up the initial state of the game.
func (g *game) initialize() {
	g.backgroundColor = color.RGBA{0, 181, 226, 255}

	// Load player image.
	g.playerImg = loadImage("asset/image/shrimp.png")

	// Reset game state.
	g.reset()

	// Initialize game attributes.
	g.fullscreen = true
	g.initialized = true
}

// Layout is called each frame with the current size of the window.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Use fixed screen size.
	return screenW, screenH
}

// Update is where we update the game state.
func (g *game) Update() error {
	// Handle initialization.
	if !g.initialized {
		g.initialize()
	}

	// Handle fullscreen.
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
		g.fullscreen = !g.fullscreen
		ebiten.SetFullscreen(g.fullscreen)
	}

	// Apply gravity.
	maxFallSpeed := 0.8
	if g.playerVelocity < maxFallSpeed {
		g.playerVelocity += 0.015
		if g.playerVelocity > maxFallSpeed {
			g.playerVelocity = maxFallSpeed
		}
	}

	// Handle jump.
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.playerVelocity = -1
	}

	// Apply velocity.
	g.playerY += g.playerVelocity

	// Handle screen edge collision.
	if g.playerY < 0 || g.playerY >= screenH-playerH {
		g.reset()
	}
	return nil
}

// Draw is where we draw the game state.
func (g *game) Draw(screen *ebiten.Image) {
	// Draw background.
	screen.Fill(g.backgroundColor)

	// Draw player.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX, g.playerY)
	screen.DrawImage(g.playerImg, op)
}
