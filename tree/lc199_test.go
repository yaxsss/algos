package tree

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	cases := []struct {
		arr      []interface{}
		expected []int
	}{
		{[]interface{}{1, 2, 3, nil, 5, nil, 4}, []int{1, 3, 4}},
		{[]interface{}{1, nil, 3}, []int{1, 3}},
		{[]interface{}{}, []int{}},
	}
	for _, c := range cases {
		head := MakeTree(c.arr, Hierarchical)
		res := rightSideView(head)
		if !reflect.DeepEqual(res, c.expected) {
			t.Fatalf("%v not equal to %v", res, c.expected)
		}
	}
}
