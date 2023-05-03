package sorting

// InsertionSort sorts the given list using the insertion sort algorithm.
func InsertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		MoveLeft(list, i)
	}
}
