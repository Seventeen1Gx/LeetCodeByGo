package LeetCode

import "strings"

func ValidParentheses_1(s string) bool {
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

func ValidParentheses_2(s string) bool {
	hash := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	n := len(s)
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			// 遇到左括号，将其对应右括号入栈
			stack = append(stack, hash[s[i]])
		} else {
			// 遇到右括号，判断栈顶是否和当前相同
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
