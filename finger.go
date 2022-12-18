package main

func (finger Finger) touch(field Field, where [2]int) float64 {
	return finger.action * field.at(whereTargeted(finger.target, where))
}

func whereTargeted(target [2]int, where [2]int) [2]int {
	return [2]int{target[0] + where[0], target[1] + where[1]}
}
