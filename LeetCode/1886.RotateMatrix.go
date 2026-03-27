package LeetCode

func RotateMatrix(mat [][]int, target [][]int) bool {
	// 顺时针 90 度旋转：第 i 行变成 n-1-i 列，第 j 列变成第 j 行
	n := len(mat)
	ans1, ans2, ans3, ans4 := true, true, true, true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 不旋转
			if ans1 && mat[i][j] != target[i][j] {
				ans1 = false
			}
			// 旋转一次
			if ans2 && mat[j][n-1-i] != target[i][j] {
				ans2 = false
			}
			// 旋转二次
			if ans3 && mat[n-1-i][n-1-j] != target[i][j] {
				ans3 = false
			}
			// 旋转三次
			if ans4 && mat[n-1-j][i] != target[i][j] {
				ans4 = false
			}
		}
	}

	return ans1 || ans2 || ans3 || ans4
}
