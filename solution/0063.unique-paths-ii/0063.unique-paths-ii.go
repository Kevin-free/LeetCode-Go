package _063_unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// m : 行数；n : 列数
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// 如果第一个就有障碍
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// 初始化第一个
	obstacleGrid[0][0] = 1

	// 初始化第一列
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		}else {
			obstacleGrid[i][0] = 0
		}
	}
	// 初始化第一行
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 && obstacleGrid[0][j-1] == 1 {
			obstacleGrid[0][j] = 1
		}else{
			obstacleGrid[0][j] = 0
		}
	}
	for i := 1; i < m; i ++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}