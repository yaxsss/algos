package linklist

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
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
