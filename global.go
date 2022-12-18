package main

var fps = 60
var rows = 100
var cols = 200
var pixelHeight = 10
var pixelWidth = 5
var span = 5
var power = .8
var randomHands = 3
var integrity = 0.7
var hands = rectHands(randomHands)
var field = [2]Field{grid[float64](), grid[float64]()}
var readIndex = 0
var writeIndex = 1
