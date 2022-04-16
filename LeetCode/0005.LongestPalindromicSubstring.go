package LeetCode

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成

func longestPalindrome1(s string) string {
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

func LongestPalindrome2(s string) string {
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
