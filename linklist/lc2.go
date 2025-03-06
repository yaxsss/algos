package linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cflags := 0
	var header, prev *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}
		leftOp, rightOp := 0, 0
		if l1 != nil {
			leftOp = l1.Val
		}
		if l2 != nil {
			rightOp = l2.Val
		}
		newNode := new(ListNode)
		newNode.Val = leftOp + rightOp + cflags
		if newNode.Val >= 10 {
			cflags = 1
			newNode.Val = newNode.Val - 10
		} else {
			cflags = 0
		}
		if prev == nil {
			header = newNode
		} else {
			prev.Next = newNode
		}

		prev = newNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if cflags == 1 {
		prev.Next = &ListNode{1, nil}
	}
	return header
}
