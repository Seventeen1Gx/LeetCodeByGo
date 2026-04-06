package LeetCode

func SolveSudoku(board [][]byte) {
	var rowMap, colMap [9][9]bool
	var gridMap [3][3][9]bool

	for i, row := range board {
		for j, v := range row {
			if board[i][j] == '.' {
				continue
			}
			num := v - '0'
			if rowMap[i][num-1] || colMap[j][num-1] || gridMap[i/3][j/3][num-1] {
				// 已经不合法
				return
			}
			rowMap[i][num-1] = true
			colMap[j][num-1] = true
			gridMap[i/3][j/3][num-1] = true
		}
	}

	// 从 i j 位置开始摆放数字，是否最终有解
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == 9 {
			return true
		}
		if board[i][j] != '.' {
			if j == 8 {
				return dfs(i+1, 0)
			} else {
				return dfs(i, j+1)
			}
		}
		// 当前位置填写数字
		for num := 1; num <= 9; num++ {
			// 判断能不能填
			if rowMap[i][num-1] || colMap[j][num-1] || gridMap[i/3][j/3][num-1] {
				continue
			}
			rowMap[i][num-1] = true
			colMap[j][num-1] = true
			gridMap[i/3][j/3][num-1] = true
			board[i][j] = byte('0' + num)

			res := false
			if j == 8 {
				res = dfs(i+1, 0)
			} else {
				res = dfs(i, j+1)
			}
			if res {
				return true
			}
			// 恢复
			rowMap[i][num-1] = false
			colMap[j][num-1] = false
			gridMap[i/3][j/3][num-1] = false
			board[i][j] = '.'
		}
		return false
	}
	_ = dfs(0, 0)
}
