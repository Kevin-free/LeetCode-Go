package _22_best_time_to_buy_and_sell_stock_ii

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	dp := make([][2]int, length)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[length-1][0]

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
