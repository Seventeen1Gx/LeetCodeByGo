package LeetCode

func CountGoodPairs(nums []int) int {
	var ans int
	var cnt [101]int
	for _, num := range nums {
		ans += cnt[num]
		cnt[num]++
	}
	return ans
}
