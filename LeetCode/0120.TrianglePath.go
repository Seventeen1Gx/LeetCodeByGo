package LeetCode

import "math"

func TrianglePath_1(triangle [][]int) int {
	// 第 i 行下标最大值是 i
	m := len(triangle)
	if m == 0 {
		return 0
	}

	ans := math.MaxInt
	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {
		sum += triangle[i][j]
		if i == m-1 {
			ans = min(ans, sum)
			return
		}
		dfs(i+1, j, sum)
		if j <= i {
			dfs(i+1, j+1, sum)
		}
	}
	dfs(0, 0, 0)

	return ans
}

func TrianglePath_2(triangle [][]int) int {
	// 第 i 行的下标最大值是 i
	m := len(triangle)
	if m == 0 {
		return 0
	}
	if m == 1 {
		return triangle[0][0]
	}

	// dp[i][j] 表示到达 i,j 这个位置时候的路径和
	// dp[0][0] = triangle[0][0]
	// dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + triangle[0][0]

	ans := math.MaxInt
	dp := make([][]int, m)
	dp[0] = make([]int, 1)
	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
			if i == m-1 {
				ans = min(ans, dp[i][j])
			}
		}
	}

	return ans
}
