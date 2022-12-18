package main

var fps = 60
var rows = 100
var cols = 200
var pixelHeight = 10
var pixelWidth = 5
var span = 5
var power = Real(.8)
var randomHands = 3
var integrity = Real(0.7)
var workFieldIndex = 0
var showFieldIndex = 0
var hands = randomRectangularHands(randomHands)
var field = [2]Field{grid[Real](), grid[Real]()}
var readIndex = 0
var writeIndex = 1

// var hands = 10
// var filterings = 1
// var flickers = 1
// var useColor = false
// var filename string
