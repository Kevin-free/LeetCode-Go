package _309_best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	length := len(prices)
	// 特殊判断
	if length <= 1 {
		return 0
	}
	// 声明dp
	dp := make([][3]int, length)
	// 初始化
	dp[0][0] = 0          // 【0天】【不持股】
	dp[0][1] = -prices[0] // 【0天】【持股】
	dp[0][2] = 0          // 【0天】【冷冻期】

	for i := 1; i < length; i++ {
		// 【第i天】【不持股】 = max（昨天不持股今天不操作，昨天持股+今天卖一股）
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 【第i天】【持股】 = max（昨天持股今天不操作，昨天冷冻期+今天买一股）
		dp[i][1] = max(dp[i-1][1], dp[i-1][2]-prices[i])
		// 【第i天】【冷冻期】 = 昨天卖了不持股
		dp[i][2] = dp[i-1][0]
	}

	// 返回不持股和冷冻期的最大者
	return max(dp[length-1][0], dp[length-1][0])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
