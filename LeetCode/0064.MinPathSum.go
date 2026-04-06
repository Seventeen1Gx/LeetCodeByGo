package LeetCode

import "math"

func MinPathSum(grid [][]int) int {
	// dp[i][j] 表示到达 i,j 点的最小路径和
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			a := math.MaxInt
			if i != 0 {
				a = dp[i-1][j] + grid[i][j]
			}
			b := math.MaxInt
			if j != 0 {
				b = dp[i][j-1] + grid[i][j]
			}
			dp[i][j] = min(a, b)
		}
	}

	return dp[m-1][n-1]
}
