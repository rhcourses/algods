package code

// SearchMin returns the minimum value in a list of integers.
func SearchMin(list []int) int {
	min := list[0]
	for _, v := range list {
		if v < min {
			min = v
		}
	}
	return min
}
