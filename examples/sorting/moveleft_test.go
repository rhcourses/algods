package sorting

import "testing"

func TestMoveLef(t *testing.T) {
	list := []int{3, 2, 1}
	MoveLeft(list, 2)
	if list[0] != 1 || list[1] != 3 || list[2] != 2 {
		t.Errorf("MoveLeft(%v, 2) = %v, want %v", []int{3, 2, 1}, list, []int{1, 3, 2})
	}
}
