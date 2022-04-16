package LeetCode

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s 的，而不是部分字符串。

// 1 <= s.length <= 20
// 1 <= p.length <= 30
// s 可能为空，且只包含从 a-z 的小写字母。
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *
// 保证每次出现字符 * 时，前面都匹配到有效的字符

func isMatch1(s string, p string) bool {
	// 递归法
	return isMatchRecursion(s, p, 0, 0)
}

// isMatchRecursion 判断 s[i:] 和 p[j:] 是否匹配
func isMatchRecursion(s string, p string, i, j int) bool {
	// 时刻注意数组越界问题

	if j == len(p) {
		// 空串与空串匹配
		// p 空串，s 不空串，则不匹配
		// s 空串，p 不空串，还可以作判断
		return i == len(s)
	}

	// j 肯定小于 len(p)，即 p[j] 不会越界
	firstMatch := i < len(s) && isMatchSingleChar(s[i], p[j])
	if j+1 < len(p) && p[j+1] == '*' {
		// 匹配零个或多个
		return isMatchRecursion(s, p, i, j+2) || firstMatch && isMatchRecursion(s, p, i+1, j)
	}
	return firstMatch && isMatchRecursion(s, p, i+1, j+1)
}

func isMatchSingleChar(a, b uint8) bool {
	return a == b || b == '.'
}

var dp [][]int

func isMatch2(s string, p string) bool {
	// 动态规划模拟上述递归过程
	// dp[i][j] 表示 s[i:] 和 p[j:] 的匹配情况
	var m, n = len(s), len(p)
	dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			dp[i][j] = -1
		}
	}

	return dpT2B(s, p, 0, 0)
}

func dpT2B(s string, p string, i, j int) bool {
	// 自顶向下
	if dp[i][j] != -1 {
		// 如果有结果，直接返回结果，相当于递归剪枝，减少重复计算
		return dp[i][j] == 1
	}

	var ans bool
	if len(p) == j {
		// p 空串，看 s 是否为空串
		ans = len(s) == i
	} else {
		firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			ans = dpT2B(s, p, i, j+2) || firstMatch && dpT2B(s, p, i+1, j)
		} else {
			ans = firstMatch && dpT2B(s, p, i+1, j+1)
		}
	}

	if ans {
		dp[i][j] = 1
	} else {
		dp[i][j] = 0
	}
	return ans
}

func isMatch3(s string, p string) bool {
	// 动态规划，自底向上
	// dp[i][j] 表示 s[i:] 和 p[j:] 的匹配情况

	var m, n = len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 空串匹配
	dp[m][n] = true

	// 遍历
	for i := m; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- { // dp[0:m-1][n] 即 p 空串，s 不为空串时，默认已经是 false 了，所以这里从 n-1 开始
			firstMatch := i < m && (s[i] == p[j] || p[j] == '.')
			if j+1 < n && p[j] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}
