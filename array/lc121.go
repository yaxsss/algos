package array

func maxProfit(prices []int) int {
	// minBuy存放最低买入的价格
	minBuy := prices[0]
	// maxP为最大收益，根据当前卖出的价格-之前最低买入的价格差来刷新最大收益
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minBuy { // 刷新最低买入价格
			minBuy = prices[i]
		} else if prices[i]-minBuy > maxP {
			maxP = prices[i] - minBuy
		}
	}
	return maxP
}
