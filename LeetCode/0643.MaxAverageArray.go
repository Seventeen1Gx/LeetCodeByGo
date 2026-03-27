package LeetCode

import "math"

func MaxAverage(nums []int, k int) float64 {
	windowSum := 0
	maxSum := math.MinInt
	n := len(nums)
	for right := 0; right < n; right++ {
		// 入
		windowSum += nums[right]
		left := right - k + 1
		if left < 0 {
			continue
		}

		maxSum = max(maxSum, windowSum)

		// 出
		windowSum -= nums[left]
	}

	return float64(maxSum) / float64(k)
}
