// 2022-03-04 UCloud 技术一面 算法题1

// 必须定义一个 包名为 `main` 的包，并实现 `main()` 函数。
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。


示例 1：

输入：s = "()"
输出：true

示例 2：

输入：s = "()[]{}"
输出：true

示例 3：

输入：s = "(]"
输出：false

示例 4：

输入：s = "([)]"
输出：false

示例 5：

输入：s = "{[]}"
输出：true
*/

package main

import "fmt"

// 35 - 42
// cost 7 min
func main() {
	ss := []string{
		"()", "()[]{}", "(]",
	}
	for _, s := range ss {
		isValid := isValid2(s)
		fmt.Printf("%v : %v\n", s, isValid)
	}
}

// 26:00 - 34:45
// cost 8 min
func isValid2(s string) bool {
	n := len(s)
	// special case: 奇数个，肯定不符合
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 若栈中无元素 或 和栈顶元素不匹配，返回 false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 元素入栈
			stack = append(stack, s[i])
		}
	}
	// 栈中元素是否匹配完成
	return len(stack) == 0
}
