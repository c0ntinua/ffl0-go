package main

import "image/color"

type Field [][]float64

type Color color.RGBA

type Finger struct {
	target [2]int
	action float64
}

type Hand []Finger
