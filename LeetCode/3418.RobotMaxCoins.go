package LeetCode

import "math"

func RobotMaxCoins(coins [][]int) int {
	// dp[i][j][k] 表示到达 i,j 点已经使用了 k 次感化机会时的最大金币数
	// dp[i][j][k] = max(dp[i-1][j][k]+coins[i][j], dp[i][j-1][k]+coins[i][j], dp[i-1][j][k-1], dp[i][j-1][k-1])
	m, n := len(coins), len(coins[0])

	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 3)
		}
	}
	dp[0][0][0] = coins[0][0]
	// dp[0][0][1] = 0
	// dp[0][0][2] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			for k := 0; k < 3; k++ {
				a := math.MinInt
				b := math.MinInt
				if i != 0 {
					a = dp[i-1][j][k] + coins[i][j]
					if k != 0 {
						b = dp[i-1][j][k-1]
					}
				}

				c := math.MinInt
				d := math.MinInt
				if j != 0 {
					c = dp[i][j-1][k] + coins[i][j]
					if k != 0 {
						d = dp[i][j-1][k-1]
					}
				}

				dp[i][j][k] = max(a, b, c, d)
			}
		}
	}

	return max(dp[m-1][n-1][0], dp[m-1][n-1][1], dp[m-1][n-1][2])
}
