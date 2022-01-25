package _064_minimum_path_sum

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, col)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
