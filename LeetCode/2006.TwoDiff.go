package LeetCode

func TwoDiff(nums []int, k int) int {
	ans := 0
	count := make(map[int]int)
	for _, num := range nums {
		ans += count[num-k]
		ans += count[num+k]
		count[num]++
	}
	return ans
}
