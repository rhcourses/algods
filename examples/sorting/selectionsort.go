package sorting

// SelectionSort sorts the given list using the selection sort algorithm.
func SelectionSort(list []int) {
	for i := 0; i < len(list); i++ {
		pos := SmallestPos(list[i:]) + i
		list[i], list[pos] = list[pos], list[i]
	}
}
