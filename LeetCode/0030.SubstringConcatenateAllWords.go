package LeetCode

func SubstringConcatenateAllWords_1(s string, words []string) []int {
	// 滑动窗口
	m, n := len(words), len(words[0])

	wordsHash := make(map[string]int, 0)
	for _, v := range words {
		wordsHash[v]++
	}

	ans := []int{}
	// 确定起始字符
	for start := 0; start < n; start++ {
		overload := 0 // 多余单词数
		wordsHash2 := make(map[string]int, 0)
		for right := start; right+n <= len(s); right += n { // 一步加入一个单词
			// [right:right+n] 加入窗口
			inWord := s[right : right+n]
			if wordsHash2[inWord] == wordsHash[inWord] {
				overload++
			}
			wordsHash2[inWord]++

			left := right + n - m*n
			if left < 0 {
				continue
			}

			if overload == 0 {
				ans = append(ans, left)
			}

			outword := s[left : left+n]
			wordsHash2[outword]--
			if wordsHash2[outword] == wordsHash[outword] {
				overload--
			}
		}
	}
	return ans
}
