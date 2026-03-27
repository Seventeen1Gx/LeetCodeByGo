package LeetCode

import "LeetCodeByGo/utils"

func HitBricks_1(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	// state[x][y] > i ：表示第 i 次击打后，x,y 位置是可以确认与顶部连接的砖块
	// 这样不用每次反复初始化 state 数组
	state := make([][]int, m)
	for x := range state {
		state[x] = make([]int, n)
	}

	// 第 i 次击打后，从 x,y 出发，“感染”周围砖块
	var dfs func(x, y, i int)
	dfs = func(x, y, i int) {
		if state[x][y] > i {
			return
		}
		state[x][y] = i + 1
		for _, dir := range dirs {
			xx, yy := x+dir[0], y+dir[1]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && grid[xx][yy] == 1 {
				dfs(xx, yy, i)
			}
		}
	}

	// 每次击打砖块后，和这个砖块相连的上下左右部分可能会掉落
	result := make([]int, len(hits))
	for i, hit := range hits {
		x, y := hit[0], hit[1]
		if grid[x][y] == 0 {
			result[i] = 0
			continue
		}

		// 被击打的位置变为空白
		grid[x][y] = 0

		// 从第一行的砖块开始“感染”
		for j := 0; j < n; j++ {
			if grid[0][j] == 0 {
				continue
			}
			dfs(0, j, i)
		}

		// 检查哪些砖头掉落了
		for xx := 0; xx < m; xx++ {
			for yy := 0; yy < n; yy++ {
				if grid[xx][yy] == 1 && state[xx][yy] <= i {
					result[i]++
					grid[xx][yy] = 0
				}
			}
		}
	}

	return result
}

func HitBricks_2(grid [][]int, hits [][]int) []int {
	// 逆序处理
	// 每次添加砖块，判断与顶部连接的集合里增加了多少块砖

	m, n := len(grid), len(grid[0])

	// 在 status 上处理，而不是直接用 grid 是因为当 grid == 0 时要判断是 hit 后造成的还是本来就没砖头
	status := make([][]int, m)
	for i, row := range grid {
		status[i] = append(status[i], row...)
	}

	// 最终状态
	// 那些本来应该脱落的砖头，还存在数组中，并不影响最终结果
	// 因为恢复过程中，它们不到致使他们脱落的关键步，不会与顶部连接
	for _, hit := range hits {
		x, y := hit[0], hit[1]
		status[x][y] = 0
	}

	// 初始集合状态
	top := m * n
	uf := utils.NewUnionFind(m*n + 1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if status[i][j] == 0 {
				continue
			}
			if i == 0 {
				uf.Merge(top, i*n+j)
			}
			// 由于遍历顺序，只要和左边和上边的块合并即可
			if i > 0 && status[i-1][j] == 1 {
				uf.Merge((i-1)*n+j, i*n+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				uf.Merge(i*n+j-1, i*n+j)
			}
		}
	}

	result := make([]int, len(hits))
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for k := len(hits) - 1; k >= 0; k-- {
		i, j := hits[k][0], hits[k][1]
		if grid[i][j] == 0 {
			// 本来就没砖头
			continue
		}

		status[i][j] = 1

		size := uf.Size(top)
		if i == 0 {
			uf.Merge(top, i*n+j)
		}
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && status[x][y] == 1 {
				uf.Merge(x*n+y, i*n+j)
			}
		}

		result[k] = uf.Size(top) - size - 1
		if result[k] < 0 {
			result[k] = 0
		}
	}

	return result
}
