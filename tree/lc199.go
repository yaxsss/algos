package tree

func traverse(root *TreeNode, depth int, res *[]*TreeNode) {
	if root == nil {
		return
	}
	// 分配空间
	if len(*res) < depth {
		*res = append(*res, nil)
	}
	// 优先右节点
	if root.Right != nil {
		traverse(root.Right, depth+1, res)
	}
	// 如果右节点没有录入，再放入其他兄弟节点
	if (*res)[depth-1] == nil {
		(*res)[depth-1] = root
	}
	// 最后左节点
	if root.Left != nil {
		traverse(root.Left, depth+1, res)
	}
}

func rightSideView(root *TreeNode) []int {
	var res []*TreeNode

	traverse(root, 1, &res)

	arr := make([]int, len(res))
	for i, r := range res {
		arr[i] = r.Val
	}

	return arr
}
