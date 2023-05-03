package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	list := []int{3, 2, 1}
	InsertionSort(list)
	AssertListsEqual(t, []int{1, 2, 3}, list)
}
