package LeetCode

func StockPrice6_1(prices []int, fee int) int {
	// 第 i 天结束：
	// dp[i][0] 表示未持有股票状态下的最大利润
	// dp[i][1] 表示持有股票状态下的最大利润

	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func StockPrice6_2(prices []int, fee int) int {
	// 贪心法：把 “手续费” 算进买入成本，只要有盈利就落袋为安，遇到下跌就更新最低成本价。
	// 我们不预测未来，只做一件事：
	// - 只要今天卖出能赚钱，就卖；
	// - 只要今天比之前 “更便宜”，就把买入价更新成今天。
	n := len(prices)

	// 最大化利润的前提下，如果我们手上拥有一支股票，那他的最低买入价格是多少
	// 遍历 prices 数组
	// 如果 prices[i]+fee < buy 我们预期以 buy 买一支股票，不如以 prices[i]+fee 买一支股票
	// 故 buy = prices[i]+fee
	// 如果 prices[i] > buy 我们直接卖出股票，获得 buy-prices[i] 收益
	// 但实际这时卖出不一定是最优，未来可能遇到更高的价格，故添加反悔操作
	// 当前是为拥有一支 prices[i] 的股票，即 buy = prices[i]
	// 下一次遇到上升的股票，我们获得 prices[i+1]-prices[i] 的收益，这样等价于第 i 天没进行操作，第 i+1 天卖出
	// 其余情况 prices[i] <= buy <= prices[i]+fee 股票价格并没有低到我们放弃手上的股票，也没有高到我们卖出股票
	// 不进行任何操作
	buy := prices[0] + fee

	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i] // 虚拟买回
		}
	}

	return profit
}
