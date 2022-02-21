package main

// 富途-线上一面-算法题1
// 买卖股票的最大收益（力扣121）
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	// 最小价格
	minPrice := prices[0]
	// 最大收益
	maxProfit := 0
	for i := 1; i < length; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		// 前i天的最大收益 = max(前i-1天的最大收益， 第i天的价格 - 前i-1天中的最小价格)
		maxProfit = max(maxProfit, prices[i] - minPrice)
	}
	return maxProfit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}