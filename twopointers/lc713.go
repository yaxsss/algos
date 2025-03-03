package twopointers

// 设两个指针分别为 l,r，另外设置一个变量 \mathit{tmp} 记录 [l,r] 内所有数的乘积。最开始 l,r 都在最左面，先向右移动 r，直到第一次发现 \mathit{tmp}\geq k，这时就固定 r，右移 l，直到 \mathit{tmp}\lt k。那么对于每个 r，l 是它能延展到的左边界，由于正整数乘积的单调性，此时以 r 为右端点的满足题目条件的区间个数为 r-l+1 个
func numSubarrayProductLessThanK(nums []int, k int) int {
	sum := int64(1)
	j := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum *= int64(nums[i])
		for j <= i && sum >= int64(k) {
			sum = sum / int64(nums[j])
			j++
		}
		count += i - j + 1
	}
	return count
}
