package LeetCode

func RotateImage_1(matrix [][]int) {
	// 两次翻转等于一次顺时针 90 度旋转
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// 按对角线置换
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			// 按中线左右置换
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func RotateImage_2(matrix [][]int) {
	// 一次旋转
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			// 位置空出来
			// matrix[i][j] -> matrix[j][n-1-i] -> matrix[n-1-i][n-1-j] -> matrix[n-1-j][i]
			t := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = t
		}
	}
}
