package LeetCode

func MaxScore(s string) int {
	n := len(s)
	left := make([]int, n)  // left[i] 表示 s[0:i] 中 0 的数量
	right := make([]int, n) // right[i] 表示 s[i:] 中 1 的数量

	if s[0] == '0' {
		left[0] = 1
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			left[i] = left[i-1] + 1
		} else {
			left[i] = left[i-1]
		}
	}

	if s[n-1] == '1' {
		right[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] == '1' {
			right[i] = right[i+1] + 1
		} else {
			right[i] = right[i+1]
		}
	}

	var ans int
	for i := 0; i < n-1; i++ {
		if left[i]+right[i+1] > ans {
			ans = left[i] + right[i+1]
		}
	}

	return ans
}
