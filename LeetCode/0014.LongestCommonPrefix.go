package LeetCode

func LongestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}

	helper := func(s1, s2 string) string {
		for i := 0; i < len(s1) && i < len(s2); i++ {
			if s1[i] != s2[i] {
				return s1[:i]
			}
		}
		if len(s1) < len(s2) {
			return s1
		}
		return s2
	}

	ans := strs[0]
	for i := 1; i < len(strs); i++ {
		ans = helper(ans, strs[i])
	}
	return ans
}
