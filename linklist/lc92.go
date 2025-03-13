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
			// 逆序处理完成
			if count == right {
				if newHead == nil { // newHead还没有更新，直接就是reverseHead
					newHead = reverseHead
				} else { // 否则就将reverseHead服务正常顺序的Next
					newTail.Next = reverseHead
				}
				newTail = reverseTail
			}
		} else { // 正常序部分，直接添加到新链表，维护NewHead和NewTail，
			if newHead == nil {
				newHead = node
			} else {
				newTail.Next = node
			}
			newTail = node
		}
		count++
	}
	// 最后尾巴一定要为nil
	newTail.Next = nil
	return newHead
}
