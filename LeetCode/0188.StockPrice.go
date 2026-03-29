package LeetCode

func StockPrice4(prices []int, k int) int {
	// 参考 0123 题
	n := len(prices)
	dp := make([]int, 2*k+1)
	for i := range dp {
		if i%2 == 1 {
			dp[i] = -prices[0]
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 1 {
				dp[j] = max(dp[j], dp[j-1]-prices[i])
			} else {
				dp[j] = max(dp[j], dp[j-1]+prices[i])
			}
		}
	}

	return dp[2*k]
}
