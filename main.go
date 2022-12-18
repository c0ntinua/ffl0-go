package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var canvas Canvas = grid[Color]()
var field Field = grid[Real]()
var hands = randomRectangularHands(randomHands)
var mutField = [2]Field{grid[Real](), grid[Real]()}

func main() {
	if len(os.Args) > 1 {
		loadHandsFromFile(os.Args[1])
	}
	//field.randomize()
	mutField[0].randomize()
	mutField[1].randomize()
	//field.randomize()
	rl.InitWindow(int32(cols*pixelWidth), int32(rows*pixelHeight), "ffl0")
	rl.SetTargetFPS(int32(fps))
	for !rl.WindowShouldClose() {
		respondToUser()
		// for _, hand := range hands {
		// 	field = hand.led(field)
		// }
		for _, hand := range hands {
			updateFieldWith(hand)
		}
		//canvas.updateWithField(field)
		canvas.updateWithField(mutField[showFieldIndex])
		rl.BeginDrawing()
		canvas.plot()
		rl.EndDrawing()
	}
	rl.CloseWindow()

}
