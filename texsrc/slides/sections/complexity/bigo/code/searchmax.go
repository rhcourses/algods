package code

// SearchMax returns the maximum value in a list of integers.
func SearchMax(list []int) int {
	max := list[0]
	for _, v := range list {
		if v > max {
			max = v
		}
	}
	return max
}
