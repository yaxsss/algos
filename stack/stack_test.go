package stack

import (
	"testing"
)

func TestLC71(t *testing.T) {
	cases := []struct {
		path     string
		expected string
	}{
		{"/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, c := range cases {
		p := simplifyPath(c.path)
		if p != c.expected {
			t.Fatalf("%v %v not equal to %v", c.path, p, c.expected)
		}
	}
}
func TestLC20(t *testing.T) {
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
