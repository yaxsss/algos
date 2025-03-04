package twopointers

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
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
