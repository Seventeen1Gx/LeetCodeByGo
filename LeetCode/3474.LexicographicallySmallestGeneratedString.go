package LeetCode

import "bytes"

func LexicographicallySmallestGeneratedString(str1, str2 string) string {
	n, m := len(str1), len(str2)

	// 贪心构造
	word := make([]byte, n+m-1)
	wordCopy := make([]byte, n+m-1)
	for i := 0; i < n; i++ {
		if str1[i] == 'T' {
			// 从 T 位置开始 m 个字符与 str2 相同
			for j := 0; j < m; j++ {
				if word[i+j] != 0 && word[i+j] != str2[j] {
					return ""
				}
				word[i+j] = str2[j]
				wordCopy[i+j] = str2[j]
			}
		}
	}

	// 剩下待定位置都填 'a'
	for i := 0; i < n+m-1; i++ {
		if word[i] == 0 {
			word[i] = 'a'
		}
	}

	// 检查 F 的位置要与 str2 不同
next:
	for i := 0; i < n; i++ {
		if str1[i] != 'F' {
			continue
		}
		sub := word[i : i+m]
		if !bytes.Equal(sub, []byte(str2)) {
			continue
		}
		// 在 [i:i+m) 找最后一个待定位置设置成 'b'
		for j := i + m; i >= 0; j-- {
			if wordCopy[j] == 0 {
				word[j] = 'b'
				continue next
			}
		}
		return ""
	}
	return string(word)
}
