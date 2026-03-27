package LeetCode

func BinaryMatrixFlip(grid [][]int) (ans bool) {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// 从 i,j 出发，能否到达右下角
		if i == m-1 && j == n-1 {
			return true
		}
		// 由于短路，这里只会删除最先到的那条路径（优先向下）
		grid[i][j] = 0
		return i+1 < m && grid[i+1][j] == 1 && dfs(i+1, j) ||
			j+1 < n && grid[i][j+1] == 1 && dfs(i, j+1)
	}

	ans = !dfs(0, 0)

	// 上面删除了一条路径，再看看是否还有路径可以到达
	ans = ans || !dfs(0, 0)

	return ans
}
