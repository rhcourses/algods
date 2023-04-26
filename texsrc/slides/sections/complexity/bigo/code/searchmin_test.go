package code

import "testing"

func TestSearchMin(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{5, 4, 3, 2, 1}
	list3 := []int{3, 1, 42, 5, -2}

	if SearchMin(list1) != 1 {
		t.Errorf("SearchMin(%v) != 1", list1)
	}
	if SearchMin(list2) != 1 {
		t.Errorf("SearchMin(%v) != 1", list2)
	}
	if SearchMin(list3) != -2 {
		t.Errorf("SearchMin(%v) != -2", list3)
	}
}
