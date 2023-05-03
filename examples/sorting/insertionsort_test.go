package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	list := []int{3, 2, 1}
	InsertionSort(list)
	if list[0] != 1 || list[1] != 2 || list[2] != 3 {
		t.Errorf("InsertionSort(%v) = %v, want %v", []int{3, 2, 1}, list, []int{1, 2, 3})
	}
}
