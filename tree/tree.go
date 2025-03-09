package tree

// 深度优先遍历和广度优先遍历

import "container/list"

type Order int

const (
	PreOrder Order = iota
	PostOrder
	InOrder
	Hierarchical
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(a []interface{}, order Order) *TreeNode {
	if (len(a)) == 0 || a[0] == nil {
		return nil
	}
	var header *TreeNode = &TreeNode{a[0].(int), nil, nil}
	if order == Hierarchical {
		list := &list.List{}
		list.PushBack(header)
		i := 1
		// 队列实现，递归都可以用队列替代
		for list.Len() > 0 {
			node := list.Front().Value.(*TreeNode)
			list.Remove(list.Front())
			if a[i] != nil {
				node.Left = &TreeNode{a[i].(int), nil, nil}
				list.PushBack(node.Left)
			}
			i++
			if i >= len(a) {
				break
			}
			if a[i] != nil {
				node.Right = &TreeNode{a[i].(int), nil, nil}
				list.PushBack(node.Right)
			}
			i++
			if i >= len(a) {
				break
			}
		}
	}
	return header
}
