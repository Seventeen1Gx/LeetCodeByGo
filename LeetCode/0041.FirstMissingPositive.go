package LeetCode

import "math"

func FirstMissingPositive_1(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}

	i := 0
	for {
		if !hash[i] {
			return i
		}
		i++
		if i == math.MaxInt32 {
			return math.MaxInt32 + 1
		}
	}
}

func FirstMissingPositive_2(nums []int) int {
	// 置换法：nums[i] 号学生应该坐在 nums[nums[i]-1] 上
	n := len(nums)
	for i := 0; i < n; i++ {
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 所有学生都坐在正确位置上
	return n + 1
}
