package dp
func rob(nums []int) int {
	//dp[i]表示抢劫到第i家时的最大金额
	//dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	//dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {	
		return max(nums[0], nums[1])
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[len(nums)-1]
}