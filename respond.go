package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func respondToUser() {
	switch {
	case rl.IsKeyReleased(rl.KeyQ):
		os.Exit(0)
	case rl.IsKeyReleased(rl.KeyH):
		hands = rectHands(randomHands)
	case rl.IsKeyReleased(rl.KeySpace):
		field[readIndex].randomize()
	case rl.IsKeyReleased(rl.KeyRight):
		adjustSpeed(5)
	case rl.IsKeyReleased(rl.KeyLeft):
		adjustSpeed(-5)
	case rl.IsKeyReleased(rl.KeyW):
		writeHandsToFile()
	case rl.IsKeyReleased(rl.KeyOne):
		changeHand(0)
	case rl.IsKeyReleased(rl.KeyTwo):
		changeHand(1)
	case rl.IsKeyReleased(rl.KeyThree):
		changeHand(2)
	case rl.IsKeyReleased(rl.KeyFour):
		changeHand(3)
	case rl.IsKeyReleased(rl.KeyFive):
		changeHand(4)
	case rl.IsKeyReleased(rl.KeySix):
		changeHand(5)
	case rl.IsKeyReleased(rl.KeySeven):
		changeHand(6)
	case rl.IsKeyReleased(rl.KeyEight):
		changeHand(7)
	case rl.IsKeyReleased(rl.KeyNine):
		changeHand(8)
	case rl.IsKeyReleased(rl.KeyZero):
		changeHand(9)
	}

}

func adjustSpeed(i int) {
	fps += i
	if fps <= 0 {
		fps = 1
	}
	rl.SetTargetFPS(int32(fps))
	fmt.Print("fps = ", fps)
}
