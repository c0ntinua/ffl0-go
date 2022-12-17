package main

func grid[T any]() [][]T {
	var grid = make([][]T, rows)
	for i := range grid {
		grid[i] = make([]T, cols)
	}
	return grid
}
