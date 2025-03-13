package linklist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	p := head
	// 这里有个小技巧，就是将新生成的节点直接先加入到节点后边，然后处理Random指针时很方便
	for p != nil {
		newNode := &Node{p.Val, p.Next, nil}
		p.Next = newNode
		p = p.Next.Next
	}
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	// 剥离链表，生成新链表
	newHead := head.Next
	p1 := newHead
	p = head
	for p1.Next != nil {
		p.Next = p1.Next
		p1.Next = p1.Next.Next
		p = p.Next
		p1 = p1.Next
	}
	// 最后一个节点的Next一定要为nil, 很多情况下编码都要注意
	p.Next = nil
	return newHead
}
