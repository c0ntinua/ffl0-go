package main

func (field Field) at(w [2]int) float64 {
	return field[_f(w[0], rows)][_f(w[1], cols)]
}

func (field Field) randomize() {
	for r := range field {
		for c := range field[r] {
			field[r][c] = randomReal(-1, 1)
		}
	}
}

func _f(x, m int) (y int) {
	switch {
	case 0 <= x && x < m:
		y = x
	case x > m:
		y = x - m
	case x < m:
		y = x + m
	}
	return
}
