package LeetCode

import (
	"math"
	"sort"
)

func MinAbsDiffSubmatrix(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])

	ans := make([][]int, m-k+1)
	for i := range ans {
		ans[i] = make([]int, n-k+1)
	}

	// 枚举左上角
	for i := 0; i < m-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			nums := make([]int, 0, k*k)
			for x := i; x < i+k; x++ {
				for y := j; y < j+k; y++ {
					nums = append(nums, grid[x][y])
				}
			}
			sort.Ints(nums)

			ans[i][j] = math.MaxInt
			for idx := k*k - 1; idx > 0; idx-- {
				if nums[idx] != nums[idx-1] {
					ans[i][j] = min(ans[i][j], nums[idx]-nums[idx-1])
				}
			}
			if ans[i][j] == math.MaxInt {
				ans[i][j] = 0
			}
		}
	}

	return ans
}
