package LeetCode

func Transpos(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[j][i] = matrix[i][j]
		}
	}
	return ans
}
