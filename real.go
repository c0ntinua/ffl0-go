package main

import "math/rand"

func randomReal(a, b float64) float64 {
	return a + (b-a)*rand.Float64()
}
