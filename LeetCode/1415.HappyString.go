package LeetCode

func GetHappyString_1(n int, k int) string {
	// 回溯法
	result := make([]string, 0, k)

	var backtrace func(curr string)
	backtrace = func(curr string) {
		if len(result) >= k {
			return
		}
		if len(curr) == n {
			result = append(result, curr)
			return
		}

		for _, c := range []byte{'a', 'b', 'c'} {
			if len(curr) == 0 || curr[len(curr)-1] != c {
				backtrace(curr + string(c)) // 字符串是不变常量，传下去的是新字符串，不影响当前字符串
			}
		}
	}

	backtrace("")

	if len(result) >= k {
		return result[k-1]
	}
	return ""
}

func GetHappyString_2(n int, k int) string {
	// 回溯法
	result := make([]string, 0, k)
	path := make([]byte, 0, n)

	var dfs func()
	dfs = func() {
		if len(result) >= k {
			return
		}
		if len(path) == n {
			result = append(result, string(path))
			return
		}

		for _, c := range []byte{'a', 'b', 'c'} {
			if len(path) == 0 || path[len(path)-1] != c {
				path = append(path, c) // 添加新字符
				dfs()
				path = path[:len(path)-1] // 移除添加的字符
			}
		}
	}

	dfs()

	if len(result) >= k {
		return result[k-1]
	}
	return ""
}

func GetHappyString_3(n int, k int) string {
	// 回溯法，无需记录每个结果，统计到第 k 个字符串就退出
	result := ""
	count := 0

	var backtrace func(curr string)
	backtrace = func(curr string) {
		if count >= k {
			return
		}
		if len(curr) == n {
			count++
			result = curr
			return
		}

		for _, c := range []byte{'a', 'b', 'c'} {
			if len(curr) == 0 || curr[len(curr)-1] != c {
				backtrace(curr + string(c))
			}
		}
	}

	backtrace("")

	if count >= k {
		return result
	}
	return ""
}

func GetHappyString_4(n int, k int) string {
	// 数学法：所有可能的字符串数 = 第一个位置 3 种选择 + 之后每个位置与之前字符不同 2 种选择
	total := 3 * (1 << (n - 1))
	if total < k {
		return ""
	}

	// 逐个确定每个位置的元素
	// 以开头元素不同分组，每组有 1 << (n-1) 个字符串（首位置确定，剩下位置各 2 种选择）
	// 元素 a 开头：aba abc aca acb
	// 元素 b 开头：bab ...
	// 元素 c 开头：cab ...
	// 我们能确定第 k 个字符在哪个组，就能确定开头元素
	// 然后确定同理下个字符

	k-- // 从零开始考虑
	result := make([]byte, n)
	chs := []byte{'a', 'b', 'c'}

	blockSize := 1 << (n - 1)
	firstCharIdx := k / blockSize
	k -= firstCharIdx * blockSize

	result[0] = chs[firstCharIdx]

	// 确定剩余字符
	for i := 1; i < n; i++ {
		blockSize >>= 1
		// 对于可选位置 i 所填字符是除前一个字符的另外两个字符
		for _, ch := range chs {
			if ch != result[i-1] {
				if k < blockSize {
					// 就在当前 ch 开头的组
					result[i] = ch
					break
				}
				k -= blockSize
			}
		}
	}

	return string(result)
}
