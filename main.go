package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	if len(os.Args) > 1 {
		loadHandsFromFile(os.Args[1])
	}
	field[0].randomize()
	field[1].randomize()
	rl.InitWindow(int32(cols*pixelWidth), int32(rows*pixelHeight), "ffl0")
	rl.SetTargetFPS(int32(fps))
	for !rl.WindowShouldClose() {
		respondToUser()
		for _, hand := range hands {
			apply(hand)

		}
		rl.BeginDrawing()
		plot(field[readIndex])
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
