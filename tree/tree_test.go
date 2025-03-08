package tree

import (
	"reflect"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	cases := []struct {
		arr      []interface{}
		expected int
	}{
		{[]interface{}{2, 9, 20, nil, nil, 15, 7}, 3},
		{[]interface{}{0, nil, 2}, 2},
	}

	for _, c := range cases {
		tree := MakeTree(c.arr, Hierarchical)
		depth := maxDepth(tree)
		if depth != c.expected {
			t.Fatalf("%v depth %v not equal to %v", c.arr, depth, c.expected)
		}
	}
}

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

func TestIsSameTree(t *testing.T) {
	cases := []struct {
		arr1     []interface{}
		arr2     []interface{}
		expected bool
	}{
		{[]interface{}{1, 2}, []interface{}{1, nil, 2}, false},
	}

	for _, c := range cases {
		tree1 := MakeTree(c.arr1, Hierarchical)
		tree2 := MakeTree(c.arr2, Hierarchical)
		if isSameTree(tree1, tree2) == c.expected {
			t.Fatalf("%v %v result not equal to %v", c.arr1, c.arr2, c.expected)
		}
	}

}
