package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	health   uint16
	name     string
	position rl.Vector2
	speed    float32
}

func NewPlayer(health uint16, name string, position rl.Vector2, speed float32) *Player {
	return &Player{
		health:   health,
		name:     name,
		position: position,
		speed:    speed,
	}
}
