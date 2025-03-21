package tree

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val { // 左节点为空或右节点为空或者左右节点值不等
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 左右节点是否对称
	return isMirror(root.Left, root.Right)
}
