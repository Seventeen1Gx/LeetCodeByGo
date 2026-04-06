package LeetCode

func CombinationSum_1(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := make([][]int, 0)

	// 当前在 candidates 的第 i 号元素，目标还剩 target，已有组合数为 cur
	// 每次抉择：选择第 i 号元素，不选第 i 号元素
	var dfs func(i, target int, cur []int)
	dfs = func(i, target int, cur []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}
		if i == n {
			return
		}

		// 选当前
		cur = append(cur, candidates[i])
		dfs(i, target-candidates[i], cur)
		cur = cur[:len(cur)-1]

		// 不选当前
		dfs(i+1, target, cur)
	}

	dfs(0, target, nil)
	return ans
}

func CombinationSum_2(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := make([][]int, 0)

	// 当前在 candidates 的第 i 号元素，目标还剩 target，已有组合数为 cur
	// 每次抉择：选择第 i~n-1 号元素
	var dfs func(i, target int, cur []int)
	dfs = func(i, target int, cur []int) {
		if target < 0 {
			return
		}
		if target == 0 {
			temp := make([]int, len(cur))
			copy(temp, cur)
			ans = append(ans, temp)
			return
		}

		for j := i; j < n; j++ {
			cur = append(cur, candidates[j])
			dfs(j, target-candidates[j], cur)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(0, target, nil)
	return ans
}
