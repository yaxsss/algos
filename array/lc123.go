package array

func maxProfit2(prices []int) int {
	rightProfits := make([]int, len(prices))
	// rightMax存放最大的卖出价格
	rightMax, right := 0, 0
	// 第i天买入时(只卖一次)的最大利益
	for i := len(prices) - 1; i >= 0; i-- {
		rightMax = max(rightMax, prices[i])
		right = max(right, rightMax-prices[i])
		rightProfits[i] = right
	}

	// leftMin用来存放最低买入时的价格
	leftMin := prices[0]
	ret := max(0, rightProfits[0])
	for i := 1; i < len(prices)-1; i++ {
		// prices[i] - leftMin是第i天卖出时的最大收益 + 第i+1天买入时(只卖一次)的最大收益
		ret = max(ret, prices[i]-leftMin+rightProfits[i+1])
		leftMin = min(leftMin, prices[i])
	}
	return ret
}
