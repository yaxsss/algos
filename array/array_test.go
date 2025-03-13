package array

import (
	"reflect"
	"testing"
)

func TestLC189(t *testing.T) {
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
func TestLC169(t *testing.T) {
	e := majorityElement([]int{3, 2, 3})
	if e != 3 {
		t.Fatalf("%d not equal 3", e)
	}
	e = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	if e != 2 {
		t.Fatalf("%d not equal 2", e)
	}
}

func TestLC88(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums, 3, nums2, 3)
	if !reflect.DeepEqual(nums, []int{1, 2, 2, 3, 5, 6}) {
		t.Fatalf("%v not equal expected [1,2,2,3,5,6]", nums)
	}

	nums = []int{1}
	nums2 = []int{}
	merge(nums, 1, nums2, 0)
	if !reflect.DeepEqual(nums, []int{1}) {
		t.Fatalf("%v not equal expected [1]", nums)
	}

	nums = []int{0}
	nums2 = []int{1}
	merge(nums, 0, nums2, 1)
	if !reflect.DeepEqual(nums, []int{1}) {
		t.Fatalf("%v not equal expected [1]", nums)
	}

}

func TestLC121(t *testing.T) {
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

func TestLC123(t *testing.T) {
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
