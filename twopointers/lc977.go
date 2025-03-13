package twopointers

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// 左右指针, 因为nums中数的绝对值是沿着Y值对称的
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
