package stack

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		expr     string
		expected bool
	}{
		// {"()", true},
		{"()[]{}", true},
		{"(]", false},
	}

	for _, c := range cases {
		if isValid(c.expr) != c.expected {
			t.Fatalf("%s is not equal %v", c.expr, c.expected)
		}
	}
}
