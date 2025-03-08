package backtrack

import (
	"reflect"
	"testing"
)

// Leetcode 17
func TestLetterCombinations(t *testing.T) {
	cases := []struct {
		digits   string
		expected []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, c := range cases {
		result := letterCombinations(c.digits)
		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("digits %v %v not equal to expected %v", c.digits, result, c.expected)
		}
	}
}

func TestCombine(t *testing.T) {
	cases := []struct {
		n        int
		k        int
		expected [][]int
	}{
		{4, 2, [][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4}}},
		{1, 1, [][]int{[]int{1}}},
	}

	for _, c := range cases {
		result := combine(c.n, c.k)
		if !reflect.DeepEqual(result, c.expected) {
			t.Fatalf("%v not equal to expected %v", result, c.expected)
		}
	}
}
