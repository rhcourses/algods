package code

import "testing"

func TestClosestPair(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{5, 4, 3, 2, 1}
	list3 := []int{3, 1, 42, 5, -2}

	if ClosestPair(list1) != 1 {
		t.Errorf("ClosestPair(%v) != 1", list1)
	}
	if ClosestPair(list2) != 1 {
		t.Errorf("ClosestPair(%v) != 1", list2)
	}
	if ClosestPair(list3) != 2 {
		t.Errorf("ClosestPair(%v) != 2", list3)
	}
}
