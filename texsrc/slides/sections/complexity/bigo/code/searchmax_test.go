package code

import "testing"

func TestSearchMax(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{5, 4, 3, 2, 1}
	list3 := []int{3, 1, 42, 5, -2}

	if SearchMax(list1) != 5 {
		t.Errorf("SearchMax(%v) != 1", list1)
	}
	if SearchMax(list2) != 5 {
		t.Errorf("SearchMax(%v) != 1", list2)
	}
	if SearchMax(list3) != 42 {
		t.Errorf("SearchMax(%v) != 42", list3)
	}
}
