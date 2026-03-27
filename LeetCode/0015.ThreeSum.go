package LeetCode

import (
	"sort"
)

func ThreeSum_1(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0)

	// 排序
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			// 排除重复情况
			continue
		}

		// 优化一：如果最小的三数之和都大于 0，那么所有三数之和都大于零
		if i+2 < n && nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 优化二：如果最大的三数之和都小于 0，那么当前的三数之和都小于零
		if i < n-2 && nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		// 首尾双指针
		j := i + 1
		k := n - 1
		for j < k {
			if j != i+1 && nums[j-1] == nums[j] {
				j++
				continue
			}
			if k != n-1 && nums[k+1] == nums[k] {
				k--
				continue
			}
			sum := nums[j] + nums[k]
			if sum == -nums[i] {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum > -nums[i] {
				k--
			} else {
				j++
			}
		}
	}

	return result
}
