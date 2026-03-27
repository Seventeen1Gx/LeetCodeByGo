package LeetCode

import (
	"math/rand"
	"time"
)

// 数组长度 2n 有一个元素重复 n 次，其他元素不重复，各出现一次

func repeatedNTimes1(nums []int) int {
	var hashSet = make(map[int]int, len(nums))

	for _, num := range nums {
		if _, ok := hashSet[num]; ok {
			// 只要重复一次就是我们要的元素
			return num
		}
		hashSet[num]++
	}

	return -1
}

func repeatedNTimes2(nums []int) int {
	// 绝对众数，至少出现 n/2+1 次，而本问题中，要求的数只出现 n/2 次，那么就排除掉第一个

	var target int
	var cnt int
	for _, num := range nums[1:] {
		if nums[0] == num {
			// 考虑第一个数就是我们要找的数的可能
			return num
		}
		// 其他情况，我们要找的数就是绝对众数
		if cnt == 0 {
			target = num
			cnt++
		} else if target == num {
			cnt++
		} else {
			cnt--
		}
	}

	return target
}

func repeatedNTimes3(nums []int) int {
	// 随机两个下标
	src := rand.NewSource(time.Now().UnixNano())
	rd := rand.New(src)

	n := len(nums)
	for {
		a := rd.Intn(n)
		b := rd.Intn(n)
		if a != b && nums[a] == nums[b] {
			return nums[a]
		}
	}
}

func repeatedNTimes4(nums []int) int {
	// 检查邻近元素
	// 从某个元素往前看 3 个元素，必然出现重复
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == nums[i-1] || i >= 2 && num == nums[i-2] || i >= 3 && num == nums[i-3] {
			return num
		}
	}
	return -1
}
