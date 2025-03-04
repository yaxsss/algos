package twopointers

import "testing"

func TestThreeSum(t *testing.T) {
	cases := []struct {
		nums  []int
		count int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, 2},
		{[]int{0, 1, 1}, 0},
		{[]int{0, 0, 0}, 1},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, 6},
	}
	for _, c := range cases {
		nums := threeSum(c.nums)
		if len(nums) != c.count {
			t.Fatalf("%v output %v", c.nums, nums)
		}
	}
}
