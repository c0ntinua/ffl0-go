package main

import "image/color"

type Canvas [][]Color
type Field [][]Real

type Real float64
type Color color.RGBA

type Where [2]int
type Target [2]int

type Finger struct {
	target Target
	action Real
}

type Hand []Finger
