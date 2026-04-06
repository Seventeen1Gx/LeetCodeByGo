package LeetCode

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成

func LongestPalindrome_1(s string) string {
	// 中心扩散法
	var ans string
	var n = len(s)

	// 奇数长度
	for i := 0; i < n; i++ {
		for step := 0; ; step++ {
			low, high := i-step, i+step
			if low >= 0 && high < n && s[low] == s[high] {
				// 说明 [low:high] 是回文
			} else {
				// 说明 [low:high] 不是回文，但上一步是 [low+1:high-1] 回文
				if high-low-1 > len(ans) {
					ans = s[low+1 : high]
				}
				break
			}
		}
	}

	// 偶数长度
	for i := 0; i < n; i++ {
		for step := 0; ; step++ {
			low, high := i-step, i+1+step
			if low >= 0 && high < n && s[low] == s[high] {
				// 说明 [low:high] 是回文
			} else {
				// 说明 [low:high] 不是回文，但上一步是 [low+1:high-1] 回文
				if high-low-1 > len(ans) {
					ans = s[low+1 : high]
				}
				break
			}
		}
	}

	return ans
}

func LongestPalindrome_2(s string) string {
	// 动态规划
	// dp[i][j] 表示 s[i:j] 是否为回文串，i<=j
	var n = len(s)
	var ans string
	var dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 初始化
	for i := 0; i < n; i++ {
		// 长度为 1 的一定是回文串
		dp[i][i] = true
		if len(ans) < 1 {
			ans = s[i : i+1]
		}
		// 长度为 2 的需要做一下判断
		if i+1 < n && s[i] == s[i+1] {
			dp[i][i+1] = true
			if len(ans) < 2 {
				ans = s[i : i+2]
			}
		}
	}

	// 按长度递增进行遍历
	for step := 3; step <= n; step++ {
		for left := 0; left < n; left++ {
			right := left + step - 1
			if right < n && dp[left+1][right-1] && s[left] == s[right] {
				dp[left][right] = true
				if step > len(ans) {
					ans = s[left : right+1]
				}
			}
		}
	}
	return ans
}

func LongestPalindrome_3(s string) string {
	// 马拉车算法
	n := len(s)
	if n < 2 {
		return s
	}

	// 构造新串，统一奇偶
	// abc -> #a#b#c#
	t := make([]byte, 0, 2*n+1)
	t = append(t, '#')
	for i := 0; i < n; i++ {
		t = append(t, s[i])
		t = append(t, '#')
	}

	// p[i] 表示以 i 为中心的回文半径
	m := len(t)
	p := make([]int, m)
	C, R := 0, 0 // 从左往右遍历，当前得到的最右回文的中心和它的右边界
	maxL, center := 0, 0

	// 遍历新串每个位置
	for i := 0; i < m; i++ {
		if i < R {
			// 已知以 C 为中心，R 为右边界的回文串
			// mirror 是关于 C 的对称位置
			// C-mirror = i-C 或者 C-i = mirror-C
			mirror := 2*C - i
			p[i] = min(R-i, p[mirror])
		}

		// 中心扩展
		left, right := i-(p[i]+1), i+(p[i]+1)
		for right < m && left >= 0 && t[left] == t[right] {
			p[i]++
			left--
			right++
		}

		// 更新最右边界
		if i+p[i] > R {
			C = i
			R = i + p[i]
		}

		if p[i] > maxL {
			maxL = p[i]
			center = i
		}
	}

	// 映射回原串
	start := (center - maxL) / 2
	return s[start : start+maxL]
}
