package LeetCode

func RemoveOuterParentheses(s string) string {
	var ans string
	var leftCnt int
	var begin int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftCnt++
		} else if s[i] == ')' {
			leftCnt--
			if leftCnt == 0 {
				ans += s[begin+1 : i]
				begin = i + 1
			}
		}
	}
	return ans
}
