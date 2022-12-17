package main

import (
	"fmt"
	"strings"
)

func (finger Finger) String() string {
	return fmt.Sprintf("%+03d %+03d %+020.15f\n", finger.target[0], finger.target[1], finger.action)
}

func fingerFrom(s string) (f Finger) {
	r := strings.NewReader(s)
	fmt.Fscanf(r, "%d %d %f", &f.target[0], &f.target[1], &f.action)
	return
}
func handsFrom(s string) (hands []Hand) {
	S := strings.Split(s, "\n\n")
	for _, s := range S {
		hands = append(hands, handFrom(s))
	}
	return

}
func handFrom(s string) (hand Hand) {
	S := strings.Split(s, "\n")
	for _, s := range S {
		hand = append(hand, fingerFrom(s))
	}
	return
}

func (hand Hand) String() string {
	var s string
	for _, finger := range hand {
		s += finger.String()
	}
	return s
}

func stringFromHands(hands []Hand) string {
	var s string
	for i, hand := range hands {
		s += hand.String()
		if i < len(hands)-1 {
			s += "\n"
		}
	}
	return s
}
