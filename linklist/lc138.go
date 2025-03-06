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

	newHead := head.Next
	p1 := newHead
	p = head
	for p1.Next != nil {
		p.Next = p1.Next
		p1.Next = p1.Next.Next
		p = p.Next
		p1 = p1.Next
	}
	p.Next = nil
	return newHead
}
