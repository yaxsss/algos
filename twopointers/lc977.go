package twopointers

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// 左右指针
func sortedSquares(nums []int) []int {
	leftIndex := 0
	rightIndex := len(nums) - 1
	result := make([]int, len(nums))
	count := len(nums) - 1
	for {
		if leftIndex > rightIndex {
			break
		}
		if abs(nums[leftIndex]) > abs(nums[rightIndex]) {
			result[count] = nums[leftIndex] * nums[leftIndex]
			leftIndex++
		} else {
			result[count] = nums[rightIndex] * nums[rightIndex]
			rightIndex--
		}
		count--
	}
	return result
}
