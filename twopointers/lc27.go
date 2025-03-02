package twopointers

func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		// fastIndex检查元素不能被覆盖时，移动slowIndex
		if nums[fastIndex] != val {
			if slowIndex != fastIndex {
				nums[slowIndex] = nums[fastIndex]
			}
			slowIndex++
		}
	}
	return slowIndex
}
