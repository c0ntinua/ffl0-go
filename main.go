package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var canvas Canvas = grid[Color]()
var field Field = grid[Real]()
var hands = randomRectangularHands(randomHands)

func main() {
	if len(os.Args) > 1 {
		loadHandsFromFile(os.Args[1])
	}
	field.randomize()
	rl.InitWindow(int32(cols*pixelWidth), int32(rows*pixelHeight), "ffl0")
	rl.SetTargetFPS(int32(fps))
	for !rl.WindowShouldClose() {
		respondToUser()
		for _, hand := range hands {
			field = hand.led(field)
		}
		canvas.updateWithField(field)
		rl.BeginDrawing()
		canvas.plot()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
