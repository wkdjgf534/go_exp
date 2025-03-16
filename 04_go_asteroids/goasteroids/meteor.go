package goasteroids

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"go-asteroids/assets"
)

const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
	numberOfSmallMeteorsFromLargeMeteor = 4
)

type Meteor struct {
	game *GameScene
	position Vector
	rotation float64
	movement Vector
	angle float64
	rotationSpeed float64
	sprite *ebiten.Image
}

func NewMeteor(baseVelocity float64, g *GameScene, index int) *Meteor {
	// Target the center of the screen.
	target := Vector {
		X: ScreenWidth/2,
		Y: ScreenHeight/2,
	}

	// Pick a random angle.
	angle := rand.Float64() * 2 * math.Pi

	// The distance from the center that meteor should spawn at. Half the width, add some arbitrary distance.
	r := ScreenWidth/2.0 + 500

	// Create the position vector, using th angle and simple math.
	pos := Vector {
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Keep the meteor moving towards the center of the screen.
	// Give it a random velocity.
	velocity := baseVelocity + rand.Float64()*1.5

	// Create the direction vector and normalize it.
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	normalizedDirection := direction.Normalize()

	// Create the movement vector.
	movement := Vector {
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	// Assign a sprite to the meteor.
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	// Create a meteor object and return it.
	m := &Meteor{
		game: g,
		position: pos,
		movement: movement,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		sprite: sprite,
		angle: angle,
	}

	return m
}
