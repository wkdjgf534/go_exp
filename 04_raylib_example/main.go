package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(800, 450, "Dungeon crawler")

	centerVector := rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)

	player := NewPlayer("TestPlayer", centerVector, 4.5)
	enemy := NewEnemy(125, "Orc", centerVector, 1.5)
	camera := NewCamera(centerVector, centerVector, 1.0)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyRight) {
			player.position.X += player.speed
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			player.position.X -= player.speed
		}
		if rl.IsKeyDown(rl.KeyUp) {
			player.position.Y -= player.speed
		}
		if rl.IsKeyDown(rl.KeyDown) {
			player.position.Y += player.speed
		}

		if rl.IsKeyPressed(rl.KeyEqual) && camera.Zoom < maxZoom {
			camera.Zoom += 0.5
		}
		if rl.IsKeyPressed(rl.KeyMinus) && camera.Zoom > minZoom {
			camera.Zoom -= 0.5
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(*camera)

		rl.DrawCircleV(player.position, camera.Zoom, rl.Maroon)
		rl.DrawCircleV(enemy.position, camera.Zoom, rl.Red)

		rl.EndMode2D()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
