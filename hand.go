package main

import (
	"math"
	"math/rand"
)

func apply(hand Hand) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			field[writeIndex][r][c] = _apply(hand, [2]int{r, c})
		}
	}
	readIndex, writeIndex = writeIndex, readIndex
}

func _apply(hand Hand, where [2]int) float64 {
	var sum float64
	for i := range hand {
		sum += hand[i].touch(field[readIndex], where)
	}
	return math.Tanh(sum)
}

func rectHand(R, C int) Hand {
	var hand Hand
	for r := -R; r <= R; r++ {
		for c := -C; c <= C; c++ {
			if randomReal(0, 1) < integrity {
				finger := Finger{[2]int{r, c}, randomReal(-power, power)}
				hand = append(hand, finger)
			}
		}
	}
	return hand
}

func rectHands(n int) (hands []Hand) {
	for i := 0; i < n; i++ {
		hands = append(hands, randomHand())
	}
	return
}

func randomHand() Hand {
	return rectHand(rand.Int()%span+1, rand.Int()%span+1)
}

func changeHand(i int) {
	if len(hands) > i {
		hands[i] = randomHand()
	}
}
