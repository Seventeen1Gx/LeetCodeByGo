package LeetCode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	closest := 0
	dist := math.MaxInt
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			if left != i+1 && nums[left-1] == nums[left] {
				left++
				continue
			}
			if right != n-1 && nums[right+1] == nums[right] {
				right--
				continue
			}
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if sum > target {
				if sum-target < dist {
					dist = sum - target
					closest = sum
				}
				right--
			} else {
				if sum-target < dist {
					dist = sum - target
					closest = sum
				}
				left++
			}
		}
	}

	return closest
}
