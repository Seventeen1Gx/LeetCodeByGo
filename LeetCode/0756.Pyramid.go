package LeetCode

func Pyramid(bottom string, allowed []string) bool {
	hash := make(map[string][]byte) // 以 allowed 前两个字母作映射

	for _, s := range allowed {
		hash[s[:2]] = append(hash[s[:2]], s[2])
	}

	var dfs func(bottom string, t []byte) bool
	dfs = func(bottom string, t []byte) bool {
		n1, n2 := len(bottom), len(t)
		if n1 == 1 {
			return true
		}
		if n1 == n2+1 {
			return dfs(string(t), nil)
		}

		ans := false
		s := bottom[n2 : n2+2]
		for _, v := range hash[s] {
			t = append(t, v)
			ans = ans || dfs(bottom, t)
			t = t[:len(t)-1]
		}
		return ans
	}

	return dfs(bottom, nil)
}
