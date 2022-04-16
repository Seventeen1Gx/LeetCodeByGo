package LeetCode

import "sort"

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出[和]为目标值 target  的那[两个]整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func twoSum1(nums []int, target int) []int {
	// 遍历法
	ret := make([]int, 2)
	var length = len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] + nums[j] == target {
				ret[0] = i
				ret[1] = j
				break
			}
		}
	}
	return ret
}

type slice struct {
	sort.IntSlice
	idx []int
}

func (s slice) Swap(i, j int) {
	// 覆盖从 sort.IntSlice 得到的 Swap 方法
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func twoSum2(nums []int, target int) []int {
	// 排序法，但要保存每个元素在原数组的下标
	var s slice
	s.IntSlice = nums
	var length = len(nums)
	s.idx = make([]int, length)
	for i := 0; i < length; i++ {
		s.idx[i] = i
	}
	sort.Sort(s)

	ret := make([]int, 2)
	var i = 0
	var j = length - 1
	for i < j {
		if s.IntSlice[i]+s.IntSlice[j] == target {
			ret[0] = s.idx[i]
			ret[1] = s.idx[j]
			break
		} else if s.IntSlice[i]+s.IntSlice[j] < target {
			i++
		} else {
			j--
		}
	}
	return ret
}

func twoSum3(nums []int, target int) []int {
	// 哈希法，边存边找，因为 a + b = c，从 a 找 c - a，也可从 b 找 c - b
	var ret = make([]int, 2)
	var hashSet = make(map[int]int, len(nums))
	for i, num := range nums {
		// 先看有没有，再存，防止用到自己，比如 5 + 5 = 10
		j, ok := hashSet[target - num]
		if ok {
			ret[0] = i
			ret[1] = j
			break
		}
		hashSet[num] = i
	}
	return ret
}
