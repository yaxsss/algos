package array

// 考虑到这个数超过一半，可以对元素计数，如果相同的加一，如果不同的减一（抵消），当为0时选择新的元素。
// 最后留下的元素肯定是大多数
func majorityElement(nums []int) int {
	count := 0
	num := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
		}
		if num == nums[i] {
			count++
		} else {
			count--
		}
	}
	return num
}
