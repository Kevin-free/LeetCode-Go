package _062_unique_paths

// 空间复杂度 O(m*n)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 空间复杂度 O(n)
func uniquePaths2(m int, n int) int {
	// 构建一个列长的一位数组（减少空间）
	dp := make([]int, n)
	// 初始化第一行每列的值为1
	for j := 0; j < n; j++ {
		dp[j] = 1
	}
	// 从第二行开始，动态计算每一列的值
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}