package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}

	dp := make([][2]int, length)
	dp[0][0] = 0          //【第0天】【未持股】
	dp[0][1] = -prices[0] //【第0天】【持股】

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // 卖股票时扣除手续费
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]

	dp_0_0 := 0          // 压缩dp表示 第0天未持股
	dp_0_1 := -prices[0] // 压缩dp表示 第0天持股

	for i := 1; i < length; i++ {
		dp_0_0 = max(dp_0_0, dp_0_1+prices[i]-fee) // 卖股票时扣除手续费
		dp_0_1 = max(dp_0_1, dp_0_0-prices[i])
	}

	return dp_0_0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
