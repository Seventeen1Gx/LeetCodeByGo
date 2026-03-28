package LeetCode

func FindStringWithLcp(lcp [][]int) string {
	// 关键性质：LCP[i][i] = n-i, LCP[i][j] = LCP[j][i]
	// 递推关系：if word[i] == word[j], lcp[i][j] = 1 + lcp[i+1][j+1]
	//			if word[i] != word[j], lpc[i][j] = 0
	// 边界：	lcp[i][j] <= n-max(i,j)

	// 贪心构造 + 合法性校验

	// LCP[i][j] > 0 则 i 和 j 的字母必须相等，否则不相等

	n := len(lcp)

	// 构造字符串
	curChar := 'a'
	word := make([]byte, n)
	for i := 0; i < n; i++ {
		if word[i] != 0 {
			continue
		}
		if curChar > 'z' {
			return ""
		}
		// 将 i 和所有 LCP[i][j] > 0 的 j 位置赋值成 curChar
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				word[j] = byte(curChar)
			}
		}
		curChar++
	}

	// 检查合法性
	for i := 0; i < n; i++ {
		if lcp[i][i] != n-i { // 对角线
			return ""
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if lcp[i][j] != lcp[j][i] { // 对称性
				return ""
			}

			if lcp[i][j] < 0 || lcp[i][j] > n-max(i, j) { // 范围
				return ""
			}

			// 递推关系
			if word[i] == word[j] &&
				i+1 < n && j+1 < n &&
				lcp[i][j] != lcp[i+1][j+1]+1 {
				return ""
			}
			if word[i] != word[j] && lcp[i][j] > 0 {
				return ""
			}
		}
	}

	return string(word)
}
