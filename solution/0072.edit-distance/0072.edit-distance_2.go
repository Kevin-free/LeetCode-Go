package _072_edit_distance

// 备忘录
var memo [][]int

func minDistance(word1 string, word2 string) int {
	// 二维切片的初始化
	memo = make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		memo[i] = make([]int, len(word2))
	}
	return dp(word1, word2, len(word1)-1, len(word2)-1)
}

// 方法二：动态规划+备忘录
func dp(word1 string, word2 string, i, j int) int {
	// base case
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	// 字符相等，啥都不做，递归下一个
	if word1[i] == word2[j] {
		memo[i][j] = dp(word1, word2, i-1, j-1)
	} else {
		memo[i][j] = min3(
			dp(word1, word2, i, j-1)+1, // 插入
			dp(word1, word2, i-1, j)+1, // 删除
			dp(word1, word2, i-1, j-1)+1) // 替换
	}
	return memo[i][j]
}

func min3(x, y, z int) int {
	if x < y {
		return min(x, z)
	} else {
		return min(y, z)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
