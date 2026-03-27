package LeetCode

func CountGoodMeals(deliciousness []int) int {
	// 两数之和，只要下标不同，无需重复利用
	maxVal := deliciousness[0]
	for _, x := range deliciousness[1:] {
		if x > maxVal {
			maxVal = x
		}
	}

	ans := 0
	hashSet := make(map[int]int)
	for _, x := range deliciousness {
		for sum := 1; sum <= maxVal*2; sum <<= 1 {
			ans += hashSet[sum-x]
		}
		hashSet[x]++
	}
	return ans % (1e9 + 7)
}
