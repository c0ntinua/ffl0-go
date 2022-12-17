package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func respondToUser() {
	if rl.IsKeyReleased(rl.KeyQ) {
		os.Exit(0)
	}
	if rl.IsKeyReleased(rl.KeySpace) {
		field.randomize()
	}
	if rl.IsKeyReleased(rl.KeyH) {
		hands = randomRectangularHands(randomHands)
	}
	if rl.IsKeyReleased(rl.KeyW) {
		writeHandsToFile()
	}
	if rl.IsKeyReleased(rl.KeyUp) {
		fps += 5
		rl.SetTargetFPS(int32(fps))
	}
	if rl.IsKeyReleased(rl.KeyDown) {
		if fps > 6 {
			fps -= 5
			rl.SetTargetFPS(int32(fps))
		}
	}
	if rl.IsKeyReleased(rl.KeyRight) {
		fps += 1
		rl.SetTargetFPS(int32(fps))
	}
	if rl.IsKeyReleased(rl.KeyLeft) {
		if fps > 1 {
			fps -= 1
			rl.SetTargetFPS(int32(fps))
		}
	}

}
