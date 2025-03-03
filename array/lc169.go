package array

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
