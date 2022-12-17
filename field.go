package main

func (field Field) at(where Where) Real {
	if where[0] < 0 {
		where[0] += rows
	}
	if where[0] >= rows {
		where[0] -= rows
	}
	if where[1] < 0 {
		where[1] += cols
	}
	if where[1] >= cols {
		where[1] -= cols
	}
	return field[where[0]][where[1]]
}

func (field Field) randomize() {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			field[r][c] = randomReal(-1, 1)
		}
	}
}
