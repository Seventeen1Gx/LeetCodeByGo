package LeetCode

func IsValidSudoku(board [][]byte) bool {
	rowMap := [9][9]bool{}
	colMap := [9][9]bool{}
	gridMap := [9][9]bool{}

	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			if rowMap[i][v-'1'] || colMap[j][v-'1'] || gridMap[i/3*3+j/3][v-'1'] {
				return false
			}
			rowMap[i][v-'1'] = true
			colMap[j][v-'1'] = true
			gridMap[i/3*3+j/3][v-'1'] = true
		}
	}

	return true
}
