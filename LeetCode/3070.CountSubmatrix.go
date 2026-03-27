package LeetCode

func CountSubmatrixLessK(grid [][]int, k int) int {
	// 枚举右下角端点
	m, n := len(grid), len(grid[0])
	ans := 0

	dp := make([][]int, m)
	for i, row := range grid {
		dp[i] = make([]int, n)
		for j, v := range row {
			if i == 0 && j == 0 {
				dp[i][j] = v
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + v
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + v
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + v
			}

			if dp[i][j] <= k {
				ans++
			}
		}
	}

	return ans
}
