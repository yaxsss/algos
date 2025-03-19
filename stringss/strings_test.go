package stringss

import "testing"

func TestLC5(t *testing.T) {
	cases := []struct {
		s        string
		expected string
	}{
		{"ababd", "aba"},
		{"cbbd", "bb"},
		{"abaabababa", "abababa"},
		{"abcddcbaabcddc", "cddcbaabcddc"},
		{"aaabbaaaa", "aaabbaaa"},
	}

	for _, c := range cases {
		ret := longestPalindrome2(c.s)
		if ret != c.expected {
			t.Fatalf("%v result %v not equal to %v", c.s, ret, c.expected)
		}
	}
}
