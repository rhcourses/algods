package sorting

// Merge merges the two sorted sublists into a single sorted list.
func Merge(list1, list2 []int) []int {
	result := []int{}

	for len(list1) > 0 && len(list2) > 0 {
		if list1[0] < list2[0] {
			result = append(result, list1[0])
			list1 = list1[1:]
		} else {
			result = append(result, list2[0])
			list2 = list2[1:]
		}
	}
	result = append(result, list1...)
	result = append(result, list2...)

	return result
}
