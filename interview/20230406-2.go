package main

// 2023-04-06 深圳夏之岛-技术一面-算法题2

// 459. 重复的子字符串

// 给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。

// 示例 1:

// 输入: s = "abab"
// 输出: true
// 解释: 可由子串 "ab" 重复两次构成。
// 示例 2:

// 输入: s = "aba"
// 输出: false
// 示例 3:

// 输入: s = "abcabcabcabc"
// 输出: true
// 解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)

// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/repeated-substring-pattern
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 前缀表（不减一）的代码实现
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	j := 0
	next := make([]int, n)
	next[0] = j
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	// next[n-1]  最长相同前后缀的长度
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}
