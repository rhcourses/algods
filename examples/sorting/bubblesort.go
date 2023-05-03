package sorting

// BubbleSort sorts the given list using the bubble sort algorithm.
func BubbleSort(list []int) {
	for i := len(list) - 1; i > 0; i-- {
		BubbleUp(list[:i+1])
	}
}
