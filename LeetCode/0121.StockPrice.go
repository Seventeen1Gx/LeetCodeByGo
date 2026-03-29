package LeetCode

import "math"

func StockPrice(prices []int) int {
	ans := 0
	leftMin := math.MaxInt
	for _, v := range prices {
		ans = max(v-leftMin, ans)
		leftMin = min(leftMin, v)
	}
	return ans
}
