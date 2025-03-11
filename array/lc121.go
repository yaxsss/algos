package array

func maxProfit(prices []int) int {
	minBuy := prices[0]
	maxP := 0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minBuy
		if profit > maxP {
			maxP = profit
		}

		if profit < 0 {
			minBuy = prices[i]
		}
	}
	return maxP
}
