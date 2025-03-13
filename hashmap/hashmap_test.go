package hashmap

import "testing"

func TestLC205(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}

	for _, c := range cases {
		res := isIsomorphic(c.s, c.t)
		if res != c.expected {
			t.Fatalf("%v, %v not return %v", c.s, c.t, c.expected)
		}
	}
}
