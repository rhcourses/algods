package code

// DiffMinMax returns the difference between the maximum and minimum values
func DiffMinMax(list []int) int {
	return SearchMax(list) - SearchMin(list)
}
