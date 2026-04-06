package LeetCode

func GenerateParenthesis(n int) []string {
	ans := make([]string, 0)

	// 枚举当前位置，左括号还是右括号
	// 当前已有 l 个左括号，r 个右括号
	var dfs func(l, r int, cur string)
	dfs = func(l, r int, cur string) {
		if r == n {
			// 右括号摆放完毕
			ans = append(ans, cur)
			return
		}

		if l < n {
			// 摆放左括号
			dfs(l+1, r, cur+"(")
		}
		if l > r {
			// 摆放右括号
			dfs(l, r+1, cur+")")
		}
	}
	dfs(0, 0, "")

	return ans
}
