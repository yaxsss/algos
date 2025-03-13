package twopointers

import (
	"reflect"
	"testing"
)

func TestLC15(t *testing.T) {
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

func TestLC11(t *testing.T) {
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

func TestLC76(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, c := range cases {
		r := minWindow(c.s, c.t)
		if r != c.expected {
			t.Fatalf("%v not to equal expected %v", r, c.expected)
		}
	}
}

func TestLC30(t *testing.T) {
	cases := []struct {
		s        string
		t        []string
		expected []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
		{"wordthregoodgoodgoodgoodthrebestword", []string{"word", "good", "best", "thre", "good", "good"}, []int{12}},
		{"barbarfoobarthebar", []string{"bar", "bar", "foo", "the"}, []int{3, 6}},
		{"lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}, []int{13}},
		{"aaaaaaaaaaaaaa", []string{"aa", "aa"}, []int{0, 2, 4, 6, 8, 10, 1, 3, 5, 7, 9}},
	}
	for _, c := range cases {
		res := findSubString(c.s, c.t)
		if !reflect.DeepEqual(res, c.expected) {
			t.Fatalf("%v not equal to %v", res, c.expected)
		}
	}
}

func TestLC125(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{",.", true},
		{"0P", false},
		{"\"Sue,\" Tom smiles, \"Selim smote us.\"", true},
	}
	for _, c := range cases {
		if isPalindrome(c.s) != c.expected {
			t.Fatalf("%v expect %v", c.s, c.expected)
		}
	}
}

func TestLC713(t *testing.T) {
	c := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	if c != 8 {
		t.Fatalf("%d not equal 8", c)
	}

	c = numSubarrayProductLessThanK([]int{1, 2, 3}, 0)
	if c != 0 {
		t.Fatalf("%d not equal 0", c)
	}
}
