package linklist

import "testing"

func constructListNode(arr []int, pos int) *ListNode {
	var header, p, cycleNode *ListNode
	for _, a := range arr {
		node := &ListNode{a, nil}
		if p != nil {
			p.Next = node
		} else {
			header = node
		}
		p = node
		pos--
		if pos == 0 {
			cycleNode = node
		}
	}
	if cycleNode != nil {
		p.Next = cycleNode
	}
	return header
}
func TestHasCycle(t *testing.T) {
	cases := []struct {
		arr      []int
		pos      int
		expected bool
	}{
		{[]int{3, 2, 0, -4}, 2, true},
	}
	for _, c := range cases {
		header := constructListNode(c.arr, c.pos)
		if hasCycle(header) != c.expected {
			t.Fatalf("%v pos: %d not equal to %v", c.arr, c.pos, c.expected)
		}
	}
}
