package sorting

// MoveLeft moves the element at the given position to the left until it is in the correct position.
func MoveLeft(list []int, pos int) {
	for pos > 0 && list[pos] < list[pos-1] {
		list[pos], list[pos-1] = list[pos-1], list[pos]
		pos--
	}
}
