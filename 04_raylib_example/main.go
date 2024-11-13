package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(800, 450, "Dungeon crawler")
	rl.HideCursor()

	centerVector := rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)

	player := NewPlayer("TestPlayer", centerVector, 1.5)
	enemy := NewEnemy("Orc", centerVector, 1.0)
	camera := NewCamera(centerVector, centerVector, 2.0)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		movePlayer(player)

		if rl.IsKeyPressed(rl.KeyEqual) && camera.Zoom < maxZoom {
			camera.Zoom += 0.5
		}
		if rl.IsKeyPressed(rl.KeyMinus) && camera.Zoom > minZoom {
			camera.Zoom -= 0.5
		}

		cursor := rl.GetMousePosition()
		cursorCamera := rl.GetScreenToWorld2D(cursor, *camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(*camera)

		rl.DrawCircleV(player.position, camera.Zoom, rl.Maroon)
		rl.DrawCircleV(enemy.position, camera.Zoom, rl.Red)
		rl.DrawCircleLines(int32(cursorCamera.X), int32(cursorCamera.Y), 4, rl.Red)

		rl.EndMode2D()

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func movePlayer(player *Player) {
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
}
