package main

func (finger Finger) touch(field Field, where Where) Real {
	return finger.action * field.at(whereTargeted(finger.target, where))
}

func whereTargeted(target Target, where Where) Where {
	return Where{target[0] + where[0], target[1] + where[1]}
}
