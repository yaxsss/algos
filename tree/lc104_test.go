package tree

import "testing"

func TestMaxDepth(t *testing.T) {
	cases := []struct {
		arr      []interface{}
		expected int
	}{
		{[]interface{}{3, 9, 20, nil, nil, 15, 7}, 3},
		{[]interface{}{1, nil, 2}, 2},
	}

	for _, c := range cases {
		tree := MakeTree(c.arr, Hierarchical)
		depth := maxDepth(tree)
		if depth != c.expected {
			t.Fatalf("%v depth %v not equal to %v", c.arr, depth, c.expected)
		}
	}
}
