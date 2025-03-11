package array

func maxProfit2(prices []int) int {
	rightProfits := make([]int, len(prices))
	rightMax, right := 0, 0
	// 第i天买入时的最大利益
	for i := len(prices) - 1; i >= 0; i-- {
		rightMax = max(rightMax, prices[i])
		right = max(right, rightMax-prices[i])
		rightProfits[i] = right
	}

	// 最大利益就是第i天卖出时 + 第i+1天买入时的利益
	leftMin := prices[0]
	ret := max(0, rightProfits[0])
	for i := 1; i < len(prices)-1; i++ {
		ret = max(ret, prices[i]-leftMin+rightProfits[i+1])
		leftMin = min(leftMin, prices[i])
	}
	return ret
}
