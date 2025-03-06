package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToArray(node *ListNode) []int {
	ans := []int{}
	for n := node; n != nil; n = n.Next {
		ans = append(ans, n.Val)
	}
	return ans
}
