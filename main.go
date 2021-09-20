package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	g := NewGame()
	g.Setup()
	g.Preload()
	for !rl.WindowShouldClose(){
		dt := rl.GetFrameTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		g.Update(dt)
		rl.EndDrawing()
	}
	g.TearDown()
}
