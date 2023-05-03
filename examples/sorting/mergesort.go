package sorting

// MergeSort sorts the given list using the merge sort algorithm.
func MergeSort(list []int) {
	if len(list) < 2 {
		return
	}

	mid := len(list) / 2
	left := list[:mid]
	right := list[mid:]

	MergeSort(left)
	MergeSort(right)

	result := Merge(left, right)
	copy(list, result)
}
