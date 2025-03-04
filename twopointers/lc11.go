package twopointers

func maxArea(height []int) int {
	head := 0
	tail := len(height) - 1
	area := 0
	// 长递减，逐渐增加高

	for head < tail {
		area = max(area, min(height[head], height[tail])*(tail-head))
		if height[head] < height[tail] {
			head++
		} else {
			tail--
		}
	}
	return area
}
