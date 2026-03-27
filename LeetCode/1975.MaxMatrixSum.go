package LeetCode

import (
	"LeetCodeByGo/utils"
	"math"
)

func MaxMatrixSum(matrix [][]int) int64 {
	// 矩阵里任意两个负数，都可以在它们之间路径两两操作的方式，改成正数
	// 那么矩阵里如果有偶数个负数，那都能变成正数
	// 如果有奇数个负数，那最终留下一个负数，留下【所有数中绝对值最小】的那个即可

	var minAbs int = math.MaxInt
	var cnt int
	var sum int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			x := matrix[i][j]
			if x < 0 {
				cnt++
				x = -x
			}
			minAbs = utils.Min(minAbs, x)
			sum += x
		}
	}

	if cnt%2 > 0 {
		sum -= minAbs * 2
	}
	return int64(sum)
}
