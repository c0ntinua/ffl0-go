package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const scale = 1.0 / 128.0

func grayFromReal(x Real) Color {
	hue := uint8(math.Trunc((float64(x) + 1) / scale))
	return Color{hue, hue, hue, 255}
}

func (canvas Canvas) plot() {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			canvas.plotPixel(r, c)
		}
	}
}

func (canvas Canvas) plotPixel(r, c int) {
	rl.DrawRectangle(
		int32(c*pixelWidth),
		int32(r*pixelHeight),
		int32(pixelWidth),
		int32(pixelHeight),
		color.RGBA(canvas[r][c]),
	)
}

func (canvas Canvas) updateWithField(field Field) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			canvas[r][c] = grayFromReal(field[r][c])
		}
	}
}
