package LeetCode

import "strings"

func ValidParentheses(s string) bool {
	stack := make([]byte, 0)
	for _, ch := range s {
		if strings.Contains("([{", string(ch)) {
			// 左括号入栈
			stack = append(stack, byte(ch))
		} else {
			// 右括号要与栈顶元素匹配
			if len(stack) == 0 {
				return false
			}
			chTop := stack[len(stack)-1]
			if !strings.Contains("([{", string(chTop)) ||
				chTop == '(' && ch != ')' ||
				chTop == '[' && ch != ']' ||
				chTop == '{' && ch != '}' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
