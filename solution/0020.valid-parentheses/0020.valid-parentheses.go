package _020_valid_parentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := make([]byte, 0)
	for c := range s {
		if c == '(' {
			// stack.push(')')
			stack = append(stack, ')')
		}else if c == '[' {
			// stack.push(']')
			stack = append(stack, ']')
		}else if c == '{' {
			// stack.push('}')
			stack = append(stack, '}')
		}else if len(stack) == 0 || stack[len(stack)-1] != byte(c) {
			return false;
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}