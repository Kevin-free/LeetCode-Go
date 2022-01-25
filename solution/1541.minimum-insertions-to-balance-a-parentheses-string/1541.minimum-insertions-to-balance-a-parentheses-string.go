package _541_minimum_insertions_to_balance_a_parentheses_string

func minInsertions(s string) int {
	n := len(s) // 字符串长度
	count := 0  // 左括号数量
	ret := 0    // 返回结果

	i := 0
	// 遍历字符串
	for i < n {
		// 如果是左括号，count+1记录为栈
		if s[i] == '(' {
			count++
			i++
		} else {
			// 判断右括号
			// 如果连续两个右括号
			if i+1 < n && s[i+1] == ')' {
				// 判断栈中有无左括号
				if count > 0 {
					// 有，消费一个左括号
					count--
				} else {
					ret++ // 无，补一个左括号
				}
				i += 2 // 两个右括号，前进两部
			} else if i < n && s[i] == ')' {
				// 判断栈中有无左括号
				if count > 0 {
					// 有，消费一个左括号
					count--
					ret += 1 // 添一个右括号
				} else {
					ret += 2 // 添一个左括号和一个右括号
				}
				i++ // 一个右括号，前进一部
			}
		}
	}
	ret += count * 2 // 返回结果= ret + 有左括号的数*2

	return ret
}
