package leetcode

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

	a := 0
	b := 0
	tmp := 0
	for i := 2; i <= len(cost); i++ {
		tmp = min(b+cost[i-1], a+cost[i-2])
		a = b
		b = tmp
	}
	return b
}
