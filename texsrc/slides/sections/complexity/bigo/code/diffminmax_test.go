package code

import "testing"

func TestDiffMinMax(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{5, 4, 3, 2, 1}
	list3 := []int{3, 1, 42, 5, -2}

	if DiffMinMax(list1) != 4 {
		t.Errorf("DiffMinMax(%v) != 4", list1)
	}
	if DiffMinMax(list2) != 4 {
		t.Errorf("DiffMinMax(%v) != 4", list2)
	}
	if DiffMinMax(list3) != 44 {
		t.Errorf("DiffMinMax(%v) != 44", list3)
	}
}
