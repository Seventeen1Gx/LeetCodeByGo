package LeetCode

func SplitMatrixEqualSum(grid [][]int) bool {
	m, n := len(grid), len(grid[0])

	sum := 0
	for _, row := range grid {
		for _, v := range row {
			sum += v
		}
	}

	// 按行划分
	pre, suf := 0, sum
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			pre += grid[i][j]
			suf -= grid[i][j]
		}
		if pre == suf {
			return true
		}
	}

	// 按列划分
	pre, suf = 0, sum
	for j := 0; j < n-1; j++ {
		for i := 0; i < m; i++ {
			pre += grid[i][j]
			suf -= grid[i][j]
		}
		if pre == suf {
			return true
		}
	}

	return false
}
