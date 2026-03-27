package LeetCode

import (
	"slices"
	"sort"
)

// 因为题目规定 n 最大是 1000000
func BalanceNum_1(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if isBalanceNum(i) {
			return i
		}
	}
	return -1
}

func isBalanceNum(x int) bool {
	if x == 0 {
		return false
	}

	cnt := [10]int{}
	for x > 0 {
		cnt[x%10]++
		x /= 10
	}

	for i, c := range cnt {
		if c != 0 && c != i {
			return false
		}
	}

	return true
}

func BalanceNum_2(n int) int {
	// 生成平衡数列表
	var L int
	var temp = n
	for temp > 0 {
		temp /= 10
		L++
	}

	for i := 0; i < 4; i++ {
		nums := getBalanceNum(L + i)
		slices.Sort(nums)
		idx := sort.SearchInts(nums, n+1)
		if idx < len(nums) {
			return nums[idx]
		}
	}

	return -1
}

// 生成 n 位的平衡数列表
func getBalanceNum(n int) []int {
	var result = make(map[int]bool)

	// 在 digits 中摆放数字以生成平衡数，目前 digits 已摆好 digits[0:pos] ，数位统计为 cnt ，将要在 pos 摆放位置
	var gen func(digits []int, cnt []int, pos int)
	gen = func(digits []int, cnt []int, pos int) {
		if len(digits) == pos {
			// 已摆放完毕，检查是否平衡
			for i, c := range cnt {
				if c != 0 && c != i {
					return
				}
			}

			// 满足平衡条件，转换成数字
			var num int
			for i := len(digits) - 1; i >= 0; i-- {
				num = num*10 + digits[i]
			}
			// 加入到结果中
			result[num] = true
			return
		}
		// 尝试放置数字 1-9
		for i := 1; i < 10; i++ {
			if cnt[i] >= i {
				continue
			}
			digits[pos] = i
			cnt[i]++
			gen(digits, cnt, pos+1)
			cnt[i]--
		}
	}

	gen(make([]int, n), make([]int, 10), 0)

	var ret []int
	for num := range result {
		ret = append(ret, num)
	}

	return ret
}
