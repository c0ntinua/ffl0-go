package main

import (
	"io/ioutil"
	"os"
	"time"
)

func panicIf(e error) {
	if e != nil {
		panic(e)
	}
}

func fileForWriting() *os.File {
	f, err := os.Create(time.Now().String() + ".txt")
	panicIf(err)
	return f
}
func writeHandsToFile() {
	f := fileForWriting()
	defer f.Close()
	_, err := f.WriteString(stringFromHands(hands))
	panicIf(err)
}

func loadHandsFromFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	panicIf(err)
	hands = handsFrom(string(content))
}
