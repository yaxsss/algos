package linklist

import (
	"reflect"
	"testing"
)

func ArrayToNode(arr [][]int) *Node {
	if len(arr) == 0 {
		return nil
	}
	var head *Node = &Node{arr[0][0], nil, nil}

	nodes := []*Node{head}
	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &Node{arr[i][0], nil, nil}
		p = p.Next
		nodes = append(nodes, p)
	}
	p = head
	for i := 0; i < len(arr); i++ {
		if arr[i][1] != -1 {
			p.Random = nodes[arr[i][1]]
		}
		p = p.Next
	}
	return head
}

func NodeToArray(head *Node) (arr [][]int) {
	if head == nil {
		return
	}
	p := head
	nodes := make(map[*Node]int)
	index := 0
	for p != nil {
		arr = append(arr, []int{p.Val, -1})
		nodes[p] = index
		index++
		p = p.Next
	}
	p = head
	for i := range arr {
		if index, ok := nodes[p.Random]; ok {
			arr[i][1] = index
		}
		p = p.Next
	}

	return
}

func TestLC2(t *testing.T) {
	cases := []struct {
		node1, node2 *ListNode
		expected     []int
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			[]int{7, 0, 8},
		}, {
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			[]int{8, 9, 9, 9, 0, 0, 0, 1},
		}, {
			&ListNode{0, nil},
			&ListNode{0, nil},
			[]int{0},
		},
	}

	for _, c := range cases {
		ans := ListToArray(addTwoNumbers(c.node1, c.node2))
		if !reflect.DeepEqual(ans, c.expected) {
			t.Fatalf("result: %v not equal expected: %v", ans, c.expected)
		}
	}
}

func TestLC92(t *testing.T) {
	cases := []struct {
		head     *ListNode
		rang     []int
		expected []int
	}{
		//{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, []int{2, 4}, []int{1, 4, 3, 2, 5}},
		{&ListNode{3, &ListNode{5, nil}}, []int{2, 2}, []int{3, 5}},
		//{&ListNode{3, &ListNode{5, nil}}, []int{1, 2}, []int{5, 3}},
	}

	for _, c := range cases {
		head := reverseBetween(c.head, c.rang[0], c.rang[1])
		arr := ListToArray(head)
		if !reflect.DeepEqual(arr, c.expected) {
			t.Fatalf("result %v not equal to %v", arr, c.expected)
		}
	}
}

func TestLC138(t *testing.T) {
	cases := []struct {
		arr [][]int
	}{
		{[][]int{[]int{7, -1}, []int{13, 0}, []int{11, 4}, []int{10, 2}, []int{1, 0}}},
	}

	for _, c := range cases {
		head := ArrayToNode(c.arr)
		copyRandomList(head)
		res := NodeToArray(head)
		if !reflect.DeepEqual(c.arr, res) {
			t.Fatalf("result %v not equal expected %v", res, c.arr)
		}

		// newHead := copyRandomList(head)

	}
}

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
func TestLC141(t *testing.T) {
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
