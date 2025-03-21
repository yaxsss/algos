package dp

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	} else if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	// dp[i]表示到达第i阶楼梯的最小花费
	// dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	// dp[0] = 0, dp[1] = 0 第一步和第二步不花费
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		// 状态转移方程：dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
		a, b = b, min(b+cost[i-1], a+cost[i-2])
	}
	return b
}
