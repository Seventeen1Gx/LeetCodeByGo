package LeetCode

func MatrixProduct_1(grid [][]int) int {
	var ans = -1
	var n, m = len(grid), len(grid[0])
	var dfs func(curX, curY, curProduct int)

	dfs = func(curX, curY, curProduct int) {
		if curProduct == 0 && ans >= 0 {
			// 因子有零就不用往后走了
			return
		}
		curProduct = curProduct * grid[curX][curY]
		if curX == n-1 && curY == m-1 {
			// 到达终点
			if curProduct >= 0 {
				ans = max(ans, curProduct)
			}
			return
		}

		// 向右或者向下
		if curX < n-1 {
			dfs(curX+1, curY, curProduct)
		}
		if curY < m-1 {
			dfs(curX, curY+1, curProduct)
		}
	}

	dfs(0, 0, 1)

	return ans % mod
}

func MatrixProduct_2(grid [][]int) int {
	// 起点终点必经之路
	var m, n = len(grid), len(grid[0])
	if grid[0][0] == 0 || grid[m-1][n-1] == 0 {
		return 0
	}

	// po[i][j] 表示到达 (i,j) 点的最大非负积，-1 表示不存在
	// ne[i][j] 表示到达 (i,j) 点的最小非正积，1 表示不存在
	// (当前值正数看自己，负数看对面)
	var po, ne = make([][]int, m), make([][]int, m)
	for i := range po {
		po[i] = make([]int, n)
		ne[i] = make([]int, n)
	}

	// 左上角初始值
	if grid[0][0] > 0 {
		po[0][0] = grid[0][0]
		ne[0][0] = 1
	} else if grid[0][0] < 0 {
		po[0][0] = -1
		ne[0][0] = grid[0][0]
	}

	// 第一列
	for i := 1; i < m; i++ {
		x := grid[i][0]
		if x == 0 {
			continue
		}
		if x > 0 {
			po[i][0] = po[i-1][0] * x
			ne[i][0] = ne[i-1][0] * x
		} else {
			po[i][0] = ne[i-1][0] * x
			ne[i][0] = po[i-1][0] * x
		}
		if po[i][0] < 0 {
			po[i][0] = -1
		}
		if ne[i][0] > 0 {
			ne[i][0] = 1
		}
	}

	// 第一行
	for j := 1; j < n; j++ {
		x := grid[0][j]
		if x == 0 {
			continue
		}
		if x > 0 {
			po[0][j] = po[0][j-1] * x
			ne[0][j] = ne[0][j-1] * x
		} else {
			po[0][j] = ne[0][j-1] * x
			ne[0][j] = po[0][j-1] * x
		}
		if po[0][j] < 0 {
			po[0][j] = -1
		}
		if ne[0][j] > 0 {
			ne[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			x := grid[i][j]
			if x == 0 {
				continue
			}
			if x > 0 {
				po[i][j] = max(po[i-1][j]*x, po[i][j-1]*x)
				ne[i][j] = min(ne[i-1][j]*x, ne[i][j-1]*x)
			} else {
				po[i][j] = max(ne[i-1][j]*x, ne[i][j-1]*x)
				ne[i][j] = min(po[i-1][j]*x, po[i][j-1]*x)
			}
			if po[i][j] < 0 {
				po[i][j] = -1
			}
			if ne[i][j] > 0 {
				ne[i][j] = 1
			}
		}
	}

	return po[m-1][n-1] % mod
}
