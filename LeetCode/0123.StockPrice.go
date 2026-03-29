package LeetCode

import "math"

func StockPrice3_1(prices []int) int {
	n := len(prices)
	ans := math.MinInt
	var dfs func(i, hold, curProfit, cnt int)
	dfs = func(i, hold, curProfit, cnt int) {
		if i == n || cnt == 2 {
			ans = max(ans, curProfit)
			return
		}
		if hold < 0 {
			// 未持有股票
			dfs(i+1, -1, curProfit, cnt)
			dfs(i+1, i, curProfit-prices[i], cnt)
		} else {
			// 持有股票
			dfs(i+1, hold, curProfit, cnt)
			dfs(i+1, -1, curProfit+prices[i], cnt+1)
		}
	}

	dfs(0, -1, 0, 0)
	return ans
}

func StockPrice3_2(prices []int) int {
	// 最多完成两笔交易
	// 任意一天结束后，有以下几种状态：
	// - 未进行过任何操作
	// - 进行过一次买入操作
	// - 进行过一次买入一次卖出操作，即完成了一次交易
	// - 进行过两次买入一次卖出操作
	// - 进行过两次买入两次卖出操作，即完成了两次交易

	// 第一个状态利润为零，无需记录，剩下状态的最大利润用 buy1，sell1，buy2，sell2 记录

	// 如果知道前一天的不同状态的最大利润，怎么推算今天的不同状态的最大利润

	// 对于 buy1 来说，可以是前一天的 buy1，也可以是当天进行买入操作
	// 即 buy1 = max(buy1', -prices[i])

	// 其他状态转移
	// sell1 = max(sell1', buy1'+prices[i])
	// buy2 = max(buy2', sell1'-prices[i])
	// sell2 = max(sell2', buy2'+prices[i])

	// 又由于在同一天进行买入卖出，对答案不会有影响故转移方程变成
	// sell1 = max(sell1', buy1+prices[i])
	// buy2 = max(buy2', sell1-prices[i])
	// sell2 = max(sell2', buy2+prices[i])

	// 在第 1 天
	// buy1 = -prices[0]
	// sell1 = 0
	// buy2 = -prices[0]
	// sell2 = 0
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}
