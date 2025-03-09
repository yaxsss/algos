package stack

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
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
