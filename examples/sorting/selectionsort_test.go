package sorting

import "testing"

func TestSelectionSort(t *testing.T) {
	list := []int{3, 2, 1}
	SelectionSort(list)
	AssertListsEqual(t, []int{1, 2, 3}, list)
}
