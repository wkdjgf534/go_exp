package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	maxZoom = 5.0
	minZoom = 1.0
)

func NewCamera(offset, target rl.Vector2, zoom float32) *rl.Camera2D {
	return &rl.Camera2D{
		Offset: offset,
		Target: target,
		Zoom:   zoom,
	}
}
