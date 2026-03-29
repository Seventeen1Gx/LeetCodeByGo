package LeetCode

func PascalsTriangle2(rowIndex int) []int {
	// 滚动数组
	// C[i][j] = C[i-1][j] + C[i-1][j-1]
	prev := []int{1}
	for i := 1; i < rowIndex+1; i++ {
		cur := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				cur[j] = 1
			} else {
				cur[j] = prev[j] + prev[j-1]
			}
		}
		prev = cur
	}
	return prev
}
