package LeetCode

import "math"

func StockPrice5_1(prices []int) int {
	// 卖出后过 1 天才能再买入（冷冻期 1 天）
	ans := math.MinInt
	n := len(prices)

	var dfs func(i, hold, curProfit int)
	dfs = func(i, hold, curProfit int) {
		if i == n {
			ans = max(ans, curProfit)
			return
		}
		if hold < 0 {
			// 未持有股票，-hold 表示上次卖出的时间
			// 保持未持有
			dfs(i+1, hold, curProfit)
			if -hold > i || -hold < i-1 {
				// 前者标志从未发生过交易，后者标志已经过冷静期
				// 在第 i 天购入股票
				dfs(i+1, i, curProfit-prices[i])
			}
		} else {
			// 持有股票
			// 保持持有
			dfs(i+1, hold, curProfit)
			// 卖出
			dfs(i+1, -i, curProfit+prices[i])
		}
	}

	dfs(0, -n, 0)
	return ans
}

func StockPrice5_2(prices []int) int {
	// 第 i 天【结束后】有三种状态：
	// - 持有一支股票
	// - 未持有股票，且当天卖出【即处于冷静期】
	// - 未持有股票，且当天未卖出【即未处于冷静期】
	//
	// dp[i][0]、dp[i][1]、dp[i][2] 对应上述三种状态下的最大利润

	// dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
	// - 第 i-1 天持有股票，第 i 天未发生卖出动作
	// - 第 i-1 天未持有股票，未处于冷静期，第 i 天买入股票
	// - 无需加上 dp[i-2][1]-prices[i]，因为这种情况在 dp[i-1][2] 包含了

	// dp[i][1] = dp[i-1][0] + prices[i]
	// - 第 i-1 天持有股票，第 i 天卖出

	// dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	// - 第 i-1 天未持有股票且未处于冷静期
	// - 第 i-1 天未只有股票且处于冷静期【第 i 天解冻】

	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}

	return max(dp[n-1][1], dp[n-1][2])
}
