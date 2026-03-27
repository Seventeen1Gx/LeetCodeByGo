package LeetCode

func ShiftMatrix(mat [][]int, k int) bool {
	mod := func(a, b int) int {
		c := a % b
		if c < 0 {
			return c + b
		}
		return c
	}

	n := len(mat[0])
	for i, row := range mat {
		for j, v := range row {
			if i%2 == 0 {
				if mat[i][mod(j-k, n)] != v {
					return false
				}
			} else {
				if mat[i][mod(j+k, n)] != v {
					return false
				}
			}
		}
	}

	return true
}
