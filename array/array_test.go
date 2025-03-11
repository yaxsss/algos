package array

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		arr      []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{2, 4, 1}, 2},
	}
	for _, c := range cases {
		max := maxProfit(c.arr)
		if max != c.expected {
			t.Fatalf("err: %v  max profit %v not equal to %v", c.arr, max, c.expected)
		}
	}
}

func TestMaxProfit2(t *testing.T) {
	cases := []struct {
		arr      []int
		expected int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 13},
		{[]int{6, 1, 3, 2, 4, 7}, 7},
	}
	for _, c := range cases {
		max := maxProfit2(c.arr)
		if max != c.expected {
			t.Fatalf("err: %v  max profit %v not equal to %v", c.arr, max, c.expected)
		}
	}
}
