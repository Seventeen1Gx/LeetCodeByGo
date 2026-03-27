package LeetCode

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 提示：
// 0 <= s.length <= 5 * 10^4
// s 由英文字母、数字、符号和空格组成

func LengthOfLongestSubstring(s string) int {
	// 双指针，滑动窗口
	var left, right, ans int
	var hashSet = make(map[uint8]int) // 记录当前统计子串中包含的字符
	var n = len(s)

	for right < n {
		newChar := s[right]
		idx, ok := hashSet[newChar]
		if ok {
			// 遭遇重复，移动左边指针到重复元素的下一个位置
			for ; left <= idx; left++ {
				delete(hashSet, s[left])
			}
		}
		hashSet[newChar] = right
		right++

		if right-left > ans {
			ans = right - left
		}
	}

	return ans
}
