package LeetCode

func CountSubmatrixXY(grid [][]byte) int {
	// 枚举矩阵右下角

	// dp[i][j][2]int
	// dp[i][j][0] 表示以 (i, j) 为右下角的矩阵中 x 的数量
	// dp[i][j][1] 表示以 (i, j) 为右下角的矩阵中 y 的数量

	// 当 i > 0 && j > 0 时
	// dp[i][j][0] = dp[i-1][j][0] + dp[i][j-1][0] - dp[i-1][j-1][0] + (grid[i][j] == 'X' ? 1 : 0)
	// dp[i][j][1] = dp[i-1][j][1] + dp[i][j-1][1] - dp[i-1][j-1][1] + (grid[i][j] == 'Y' ? 1 : 0)

	// 当 i = 0 && j = 0 时
	// dp[0][0][0] = (grid[i][j] == 'X' ? 1 : 0)
	// dp[0][0][1] = (grid[i][j] == 'Y' ? 1 : 0)

	// 当 i = 0 && j > 0 时
	// dp[0][j][0] = dp[0][j-1][0] + (grid[i][j] == 'X' ? 1 : 0)
	// dp[0][j][1] = dp[0][j-1][1] + (grid[i][j] == 'Y' ? 1 : 0)

	// 当 i > 0 && j = 0 时
	// dp[i][0][0] = dp[i-1][0][0] + (grid[i][j] == 'X' ? 1 : 0)
	// dp[i][0][1] = dp[i-1][0][1] + (grid[i][j] == 'Y' ? 1 : 0)

	m, n := len(grid), len(grid[0])
	ans := 0
	dp := make([][][2]int, m)

	for i, row := range grid {
		dp[i] = make([][2]int, n)
		for j, v := range row {
			isX, isY := 0, 0
			if v == 'X' {
				isX = 1
			} else if v == 'Y' {
				isY = 1
			}

			if i == 0 && j == 0 {
				dp[i][j][0] = isX
				dp[i][j][1] = isY
			} else if i == 0 {
				dp[i][j][0] = isX + dp[i][j-1][0]
				dp[i][j][1] = isY + dp[i][j-1][1]
			} else if j == 0 {
				dp[i][j][0] = isX + dp[i-1][j][0]
				dp[i][j][1] = isY + dp[i-1][j][1]
			} else {
				dp[i][j][0] = isX + dp[i][j-1][0] + dp[i-1][j][0] - dp[i-1][j-1][0]
				dp[i][j][1] = isY + dp[i][j-1][1] + dp[i-1][j][1] - dp[i-1][j-1][1]
			}

			if dp[i][j][0] != 0 && dp[i][j][0] == dp[i][j][1] {
				ans++
			}
		}
	}

	return ans
}
