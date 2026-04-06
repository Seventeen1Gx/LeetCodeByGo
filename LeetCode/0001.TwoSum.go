package LeetCode

import (
	"slices"
)

// 给定一个整数数组 nums 和一个整数目标值 target
// 请你在该数组中找出[和]为目标值 target 的那[两个]整数，并返回它们的数组下标
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现
// 你可以按任意顺序返回答案

func TwoSum_1(nums []int, target int) []int {
	// 遍历法
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func TwoSum_2(nums []int, target int) []int {
	// 排序后双指针法
	n := len(nums)

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums[i] - nums[j] })

	i := 0
	j := n - 1
	for i < j {
		a := nums[idx[i]]
		b := nums[idx[j]]
		if a+b == target {
			return []int{idx[i], idx[j]} // 注意这里
		} else if a+b > target {
			j--
		} else {
			i++
		}
	}

	return []int{-1, -1}
}

func TwoSum_3(nums []int, target int) []int {
	// 哈希法
	n := len(nums)
	hashSet := make(map[int]int, n)

	for j, num := range nums {
		if i, ok := hashSet[target-num]; ok {
			return []int{i, j}
		}
		hashSet[num] = j
	}

	return []int{-1, -1}
}
