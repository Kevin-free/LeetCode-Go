package _695_max_area_of_island

/**
https://leetcode-cn.com/problems/max-area-of-island/
695. 岛屿的最大面积
tag：DFS
 */

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(grid, i, j))
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, i int, j int) int {
	// 不是岛屿，返回 0
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	// 遍历过的岛屿，标记为 0（沉岛操作）
	grid[i][j] = 0
	num := 1
	num += dfs(grid,i-1,j)
	num += dfs(grid,i+1,j)
	num += dfs(grid,i,j-1)
	num += dfs(grid,i,j+1)
	return num
}

func max(x int, y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}