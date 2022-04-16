package LeetCode

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 提示：
// 0 <= s.length <= 5 * 10^4
// s 由英文字母、数字、符号和空格组成

func lengthOfLongestSubstring(s string) int {
	// 双指针，滑动窗口
	var left, right, ans int
	var hashSet = make(map[uint8]int) // 记录当前统计子串中包含的字符
	var n = len(s)

	for left < n && right <= n {
		if right == n {
			// 说明统计结束
			if right-left > ans {
				ans = right - left
			}
			break
		}

		i, ok := hashSet[s[right]]
		if ok {
			// 说明遇到一个重复字符
			if right-left > ans {
				// 先统计当前子串的长度
				ans = right - left
			}
			// 移动左边界，一直移动到重复字符的下一位，明确这里 left 不会超过 right
			for left < n && left <= i {
				delete(hashSet, s[left])
				left++
			}
		} else {
			// 说明没遇到重复字符，加入统计
			hashSet[s[right]] = right
			right++
		}
	}
	return ans
}
