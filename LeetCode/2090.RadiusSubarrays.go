package LeetCode

func GetAverage(nums []int, k int) []int {
	n := len(nums)
	windowSum := 0
	ans := make([]int, n)
	for i := 0; i < k; i++ {
		ans[i] = -1
		ans[n-1-i] = -1
	}
	for right := 0; right < n; right++ {
		windowSum += nums[right]
		left := right - 2*k
		if left < 0 {
			continue
		}
		ans[right-k] = windowSum / (2*k + 1)
		windowSum -= nums[left]
	}
	return ans
}
