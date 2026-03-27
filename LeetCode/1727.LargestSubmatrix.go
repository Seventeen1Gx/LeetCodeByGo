package LeetCode

import "sort"

func LargestMatrix(matrix [][]int) (ans int) {
	n := len(matrix[0])
	heights := make([]int, n)
	for _, row := range matrix {
		for j, v := range row {
			if v == 1 {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		// 重排 heights 将最高的几项排一起
		newHeights := make([]int, n)
		copy(newHeights, heights)
		sort.Ints(newHeights)

		for i, height := range newHeights {
			// 以 newHeights[i] 为高
			ans = max(ans, (n-i)*height)
		}
	}
	return ans
}
