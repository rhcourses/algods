package sorting

// SmallestPos returns the index of the smallest element in the given list.
func SmallestPos(list []int) int {
	smallest := 0
	for i := 1; i < len(list); i++ {
		if list[i] < list[smallest] {
			smallest = i
		}
	}
	return smallest
}
