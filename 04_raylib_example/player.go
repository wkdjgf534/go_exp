package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	maxHitPoints = 255
	minHitPoints = 15
)

type Player struct {
	health   uint16
	name     string
	position rl.Vector2
	speed    float32
}

func NewPlayer(name string, position rl.Vector2, speed float32) *Player {
	return &Player{
		health:   RandomUintInRange(minHitPoints, maxHitPoints),
		name:     name,
		position: position,
		speed:    speed,
	}
}
