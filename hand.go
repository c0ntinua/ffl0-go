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
			if randomReal(0, 1) < Real(integrity) {
				finger := Finger{Target{r, c}, randomReal(-power, power)}
				hand = append(hand, finger)
			}
		}
	}
	return hand
}

func randomRectangularHands(n int) (hands []Hand) {
	for i := 0; i < n; i++ {
		hands = append(hands, randomRectangularHand())
	}
	return
}

func randomRectangularHand() Hand {
	return rectangularHand(rand.Int()%span+1, rand.Int()%span+1)
}

func changeHand(i int) {
	if len(hands) > i {
		hands[i] = randomRectangularHand()
	}
}
