package LeetCode

// 当前左括号的个数，右括号的个数需要小于等于左括号
func LongestValidParentheses_1(s string) int {
	// 动态规划
	// dp[i] 表示以 i 结尾的字串中最长有效括号的长度
	// dp[i] = dp[i-2] + 2 if s[i-1:i] = "()"
	ans := 0
	n := len(s)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else {
			l := i - 1 - dp[i-1]
			if l >= 0 && s[l] == '(' {
				dp[i] = dp[i-1] + 2
				if l-1 >= 0 {
					dp[i] += dp[l-1]
				}
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}
