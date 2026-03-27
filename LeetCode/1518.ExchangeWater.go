package LeetCode

func ExchangeWater(numBottles, numExchange int) int {
	var ans int // 喝水数
	for numBottles >= numExchange {
		n := numBottles / numExchange
		ans += n * numExchange
		numBottles %= numExchange
		numBottles += n
	}
	if numBottles > 0 {
		ans += numBottles
	}
	return ans
}
