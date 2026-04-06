package LeetCode

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)

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
		dfs(i+1, target-candidates[i], cur)
		cur = cur[:len(cur)-1]

		// 不选的时候，跳过重复
		for i+1 < n && candidates[i] == candidates[i+1] {
			i++
		}

		// 不选当前
		dfs(i+1, target, cur)
	}

	dfs(0, target, nil)
	return ans
}
