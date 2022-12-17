package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeHandsToFile() {
	f, err := os.Create(time.Now().String() + ".txt")
	check(err)
	defer f.Close()
	n, err := f.WriteString(stringFromHands(hands))
	check(err)
	fmt.Print(n, "bytes written")
}

func loadHandsFromFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	check(err)
	hands = handsFrom(string(content))
}
