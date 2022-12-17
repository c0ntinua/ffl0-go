package main

import (
	"math"
	"math/rand"
)

func (hand Hand) led(field Field) Field {
	var handledField Field = grid[Real]()
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			handledField[r][c] = hand._led(field, Where{r, c})
		}
	}
	return handledField
}

func (hand Hand) _led(field Field, where Where) Real {
	var sum Real
	for i := range hand {
		sum += hand[i].touch(field, where)
	}
	return Real(math.Tanh(float64(sum)))
}

func rectangularHand(rowSpan, colSpan int) Hand {
	var hand Hand
	for r := -rowSpan; r <= rowSpan; r++ {
		for c := -colSpan; c <= colSpan; c++ {
			finger := Finger{Target{r, c}, randomReal(-power, power)}
			hand = append(hand, finger)
		}
	}
	return hand
}

func randomRectangularHands(n int) []Hand {
	var hands = make([]Hand, n)
	for i := range hands {
		hands[i] = rectangularHand(rand.Int()%span+1, rand.Int()%span+1)
	}
	return hands
}
