package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	health   uint16
	name     string
	position rl.Vector2
	speed    float32
}

func NewEnemy(health uint16, name string, position rl.Vector2, speed float32) *Enemy {
	return &Enemy{
		health:   RandomUintInRange(minHitPoints, maxHitPoints),
		name:     name,
		position: position,
		speed:    speed,
	}
}
