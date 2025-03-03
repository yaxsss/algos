package twopointers

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	c := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	if c != 8 {
		t.Fatalf("%d not equal 8", c)
	}

	c = numSubarrayProductLessThanK([]int{1, 2, 3}, 0)
	if c != 0 {
		t.Fatalf("%d not equal 0", c)
	}
}
