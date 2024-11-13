package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(800, 450, "Dungeon crawler")

	player := NewPlayer(255, "TestPlayer", rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2), 4.5)
	enemy := NewEnemy(125, "Orc", rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2), 1.5)

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

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircleV(player.position, 2, rl.Maroon)
		rl.DrawCircleV(enemy.position, 5, rl.Red)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
