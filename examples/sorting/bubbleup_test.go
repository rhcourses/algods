package sorting

import (
	"fmt"
	"testing"
)

func TestBubbleUp(t *testing.T) {
	list := []int{1, 4, 3, 2, 6, 5}

	swapped := BubbleUp(list)
	AssertTrue(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 4, 3, 2, 6, 5}))
	AssertListsEqual(t, []int{1, 3, 2, 4, 5, 6}, list)

	swapped = BubbleUp(list)
	AssertTrue(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 3, 2, 4, 5, 6}))
	AssertListsEqual(t, []int{1, 2, 3, 4, 5, 6}, list)

	swapped = BubbleUp(list)
	AssertFalse(t, swapped, fmt.Sprintf("BubbleUp(%v)", []int{1, 2, 3, 4, 5, 6}))
}
