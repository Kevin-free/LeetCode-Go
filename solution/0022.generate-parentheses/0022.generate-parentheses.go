package _022_generate_parentheses

// var res []string

func generateParenthesis(n int) []string {
	var res []string
	//var res = new([]string)
	// 特判
	if n == 0 {
		return res
	}
	dfs("", n, n, &res)
	return res
}

// 递归，减法
func dfs(cur string, left int, right int, res *[]string) {
	//
	if left == 0 && right == 0 {
		*res = append(*res, cur)
	}
	// 剪枝
	if left > right {
		return
	}
	if left > 0 {
		dfs(cur+"(", left-1, right, res)
	}
	if right > 0 {
		dfs(cur+")", left, right-1, res)
	}

}
