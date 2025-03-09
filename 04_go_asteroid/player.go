package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	assests "go-asteroids/assets"
)

type Player struct {
	sprite *ebiten.Image
	rotation float64
}

const rotationPerSecond = math.Pi

func NewPlayer(game *Game) *Player {
	sprite := assests.PlayerSprite

	p := &Player{
		sprite: sprite,
	}

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	// Gets a half of width and height of an object
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	// Rotate an image according to what key is pressed
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}
