package tree

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}
