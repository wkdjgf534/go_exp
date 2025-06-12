package goasteroids

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"go-asteroids/assets"
)

const exhaustSpawnOffset = -50.0 // How far from the player sprite the exhaust should appear.

// Exhaust is the type for exhaust.
type Exhaust struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
}

// NewExhaust creates a new exhaust object.
func NewExhaust(pos Vector, rotation float64) *Exhaust {
	sprite := assets.ExhaustSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos.X -= halfW
	pos.Y -= halfH

	return &Exhaust{
		position: pos,
		rotation: rotation,
		sprite:   sprite,
	}
}

// Update updates the exhaust object.
func (e *Exhaust) Draw(screen *ebiten.Image) {
	bounds := e.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(e.rotation)
	op.GeoM.Translate(halfW, halfH)
	op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.sprite, op)
}

// Draw draws the exhaust object.
func (e *Exhaust) Update() {
	speed := maxAcceleration / float64(ebiten.TPS())
	e.position.X += math.Sin(e.rotation) * speed
	e.position.Y += math.Cos(e.rotation) * -speed
}
