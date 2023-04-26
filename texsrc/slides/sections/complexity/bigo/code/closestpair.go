package code

import "math"

// ClosestPair returns the smallest difference between any two values in the list.
func ClosestPair(list []int) int {
	min := DiffMinMax(list)
	for i, v1 := range list {
		for _, v2 := range list[i+1:] {
			diff := int(math.Abs(float64(v1 - v2)))
			if diff < min {
				min = diff
			}
		}
	}
	return min
}
