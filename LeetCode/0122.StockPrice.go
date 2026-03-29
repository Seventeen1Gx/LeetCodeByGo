package LeetCode

import "math"

func StockPrice2_1(prices []int) int {
	// 未持有股票时考虑是否买入
	// 持有股票时考虑是否卖出
	n := len(prices)
	profit := math.MinInt
	var dfs func(i, hold int, curProfit int)
	dfs = func(i, hold, curProfit int) {
		if i == n {
			profit = max(profit, curProfit)
			return
		}
		if hold < 0 {
			// 未持有股票
			// 可以选择在 i 天买入，也可以不买入
			dfs(i+1, i, curProfit)
			dfs(i+1, -1, curProfit)
		} else {
			// 持有股票
			// 可以选择在 i 天卖出，也可以不卖出
			dfs(i+1, -1, curProfit+(prices[i]-prices[hold]))
			dfs(i+1, hold, curProfit)
		}
	}

	dfs(0, -1, 0)
	return profit
}

func StockPrice2_2(prices []int) int {
	// dp[i][0] 表示第 i 天时手里没股票的情况下最大利润
	// dp[i][1] 表示第 i 天时手里持有 1 只股票的情况下最大利润

	// 买入和卖出单独计算利润
	// dp[0][0] = 0, dp[0][1] = -prices[0]

	// 前一天没持有股票 + 前一天持有股票，今天卖出
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

	// 前一天持有一只股票 + 前一天没持有股票，今天买入
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func StockPrice2_3(prices []int) int {
	// 贪心
	// 寻找 x 个不相交的区间 (li, ri]
	// 使得 sum(prices[ri]-prices[li]) 最大
	// (li, ri] 区间的贡献等于 (li, li+1] (li+1,li+2] ... (ri-1, ri] 区间贡献总和
	// 因为 prices[ri]-prices[li] = prices[li+1]-prices[li] + prices[li+2] - prices[li+1] + ... + prices[ri] - prices[ri-1]

	// 问题转化为寻找 x 个长度为 1 的区间 (li,li+1]，使得 sum(a[li+1]-a[li]) 最大

	ans := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		ans += max(0, prices[i]-prices[i-1])
	}

	return ans
}
