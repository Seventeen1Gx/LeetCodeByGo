package LeetCode

func TwoStrSum(nums []string, target string) int {
	// 本题并不要求 i<j
	ans := 0
	count := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if len(num) <= len(target) && num == target[len(target)-len(num):] {
			ans += count[target[:len(target)-len(num)]]
		}
		count[num]++
	}

	count = make(map[string]int)
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if len(num) <= len(target) && num == target[len(target)-len(num):] {
			ans += count[target[:len(target)-len(num)]]
		}
		count[num]++
	}

	return ans
}
