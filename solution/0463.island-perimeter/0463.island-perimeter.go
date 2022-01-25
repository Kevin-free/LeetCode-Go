package _463_island_perimeter

/**
https://leetcode-cn.com/problems/island-perimeter/
463. 岛屿的周长
tag：DFS
 */

func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(grid, i, j)
			}
		}
	}
	return 0
}

func dfs(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 1
	}
	if grid[i][j] == 0 {
		return 1
	}
	if grid[i][j] == 2 {
		return 0
	}
	grid[i][j] = 2
	return dfs(grid, i-1, j) +
		dfs(grid, i+1, j) +
		dfs(grid, i, j-1) +
		dfs(grid, i, j+1)
}
