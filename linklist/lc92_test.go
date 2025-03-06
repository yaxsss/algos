package linklist

import (
	"reflect"
	"testing"
)

func TestReverseBetween(t *testing.T) {
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
