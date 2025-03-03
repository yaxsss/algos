package tree

func traverse(root *TreeNode, depth int, res *[]*TreeNode) {
	if root == nil {
		return
	}
	if len(*res) < depth {
		*res = append(*res, nil)
	}
	if root.Right != nil {
		traverse(root.Right, depth+1, res)
	}
	if (*res)[depth-1] == nil {
		(*res)[depth-1] = root
	}
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
