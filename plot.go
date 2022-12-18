package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func plot(field Field) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			pixel(field, r, c)
		}
	}

}

func pixel(field Field, r, c int) {
	rl.DrawRectangle(
		int32(c*pixelWidth),
		int32(r*pixelHeight),
		int32(pixelWidth),
		int32(pixelHeight),
		color.RGBA(gray(field[r][c])),
	)

}

func gray(x Real) Color {
	hue := uint8(math.Trunc(128.0 * (float64(x) + 1)))
	return Color{hue, hue, hue, 255}
}
