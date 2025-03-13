package tree

// 遍历到某层就记录在result[depth]中, 后续直接求平均值
func level(root *TreeNode, depth int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) < depth {
		*result = append(*result, []int{})
	}

	(*result)[depth-1] = append((*result)[depth-1], root.Val)

	if root.Left != nil {
		level(root.Left, depth+1, result)
	}

	if root.Right != nil {
		level(root.Right, depth+1, result)
	}
}

func averageOfLevels(root *TreeNode) []float64 {
	result := [][]int{}
	level(root, 1, &result)
	res := make([]float64, len(result))
	for i, r := range result {
		count := 0
		for _, n := range r {
			count += n
		}
		res[i] = float64(count) / float64(len(r))
	}
	return res
}
