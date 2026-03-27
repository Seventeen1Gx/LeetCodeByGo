package LeetCode

import "sort"

func DiagSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 && j != 0 {
				continue
			}
			var list []int
			k := 0
			for i+k < m && j+k < n {
				list = append(list, mat[i+k][j+k])
				k++
			}
			sort.Ints(list)
			// 回填
			for t, v := range list {
				mat[i+t][j+t] = v
			}
		}
	}

	return mat
}
