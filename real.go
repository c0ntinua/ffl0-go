package main

import "math/rand"

func randomReal(a, b Real) Real {
	return a + (b-a)*Real(rand.Float64())
}
