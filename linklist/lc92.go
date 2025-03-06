package linklist

// 组成的新的链表，可以通每个节点是前插还是后插来决定反转
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}

	count := 1
	var newHead, reverseHead, newTail, reverseTail *ListNode
	for head != nil {
		node := head
		head = head.Next
		if left <= count && count <= right {
			if reverseHead != nil {
				node.Next = reverseHead
			}
			reverseHead = node
			if count == left {
				reverseTail = node
			}
			if count == right {
				if newHead == nil {
					newHead = reverseHead
				} else {
					newTail.Next = reverseHead
				}
				newTail = reverseTail
			}
		} else {
			if newHead == nil {
				newHead = node
			} else {
				newTail.Next = node
			}
			newTail = node
		}
		count++
	}
	newTail.Next = nil
	return newHead
}
