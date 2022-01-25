package _200_number_of_islands

/**
https://leetcode-cn.com/problems/number-of-islands/
200. 岛屿数量
tag：DFS
 */

func numIslands(grid [][]byte) int {
	row, col := len(grid), len(grid[0])
	islandNum := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				islandNum++
				inflect(grid, i, j)
			}
		}
	}
	return islandNum
}

// 感染函数
func inflect(grid [][]byte, i int, j int) {
	// 越界等不符合条件的，直接return
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	// 符合条件的1小岛，被感染成2
	grid[i][j] = '2'
	// 小岛可以感染相邻的四座小岛
	inflect(grid, i-1, j) //上
	inflect(grid, i+1, j) //下
	inflect(grid, i, j-1) //左
	inflect(grid, i, j+1) //右
}
