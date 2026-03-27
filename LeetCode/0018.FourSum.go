package LeetCode

import "sort"

func FourSum(nums []int, target int) [][]int {
	n := len(nums)
	result := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != i+1 && nums[j-1] == nums[j] {
				continue
			}

			k := j + 1
			l := n - 1
			for k < l {
				if k != j+1 && nums[k-1] == nums[k] {
					k++
					continue
				}
				if l != n-1 && nums[l+1] == nums[l] {
					l--
					continue
				}

				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
				} else if sum > target {
					l--
				} else {
					k++
				}
			}
		}
	}

	return result
}
