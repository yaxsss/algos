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

func TestCopyRandomList(t *testing.T) {
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
