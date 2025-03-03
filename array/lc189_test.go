package array

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		arr      []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{-1}, 2, []int{-1}},
	}
	for _, c := range cases {
		rotate(c.arr, c.k)
		if !reflect.DeepEqual(c.arr, c.expected) {
			t.Fatalf("%v not equal expected: %v", c.arr, c.expected)
		}
	}
}
