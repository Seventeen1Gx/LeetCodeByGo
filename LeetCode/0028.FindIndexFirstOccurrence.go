package LeetCode

func FindIndexFirstOccurrence_1(haystack string, needle string) int {
outer:
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue outer // 匹配失败回退两个指针
			}
		}
		return i
	}
	return -1
}

func FindIndexFirstOccurrence_2(haystack string, needle string) int {
	// KMP
	// 匹配失败只回退模式串指针，文本串指针不回退
	// 而模式串指针回退到 next 位置
	// next[j] 表示模式串前 j 个字符即 [0:j-1] 中，最长相等的前缀 & 后缀的长度
	// 即 next[j] = 最大的 k 使得 p[0..k-1] == p[j-k...j-1] 相等
	// 比如 “abaab” 中 next[5] = 2

	n, m := len(haystack), len(needle)

	next := make([]int, m+1)
	next[0] = -1 // 第一个字符匹配失败时，从头开始
	// i 用来遍历模式串
	// j 表示当前最长相等前后串
	for i, j := 0, -1; i < m; {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			// 主串不回退，模式串回退
			j = next[j]
		}
	}

	i, j := 0, 0
	for i < n && j < m {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}
