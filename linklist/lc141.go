package linklist

// hashmap
func hasCycle(head *ListNode) bool {
	// nodes := map[*ListNode]struct{}{}
	// for l := head; l != nil; l = l.Next {
	// 	if _, ok := nodes[l]; ok {
	// 		return true
	// 	}
	// 	nodes[l] = struct{}{}
	// }
	// return false

	// 题解里面提到的，快慢指针，其实环的话，N和2N一定会在某个节点上相遇
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
