package twopointers

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, c := range cases {
		area := maxArea(c.height)
		if area != c.expected {
			t.Fatalf("%v max area %v not equal to %v", c.height, area, c.expected)
		}
	}
}
