package LeetCode

import "LeetCodeByGo/utils"

func NumIslands_1(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// 引入并查集
	uf := utils.NewUnionFind(m * n)

	// 上下左右枚举
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var waterCnt int
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == '1' {
				v := x*n + y
				// 当前陆地格子与上下左右的陆地格子相连（合并到一个集合）
				for _, dir := range dirs {
					x1 := x + dir[0]
					y1 := y + dir[1]
					if x1 > 0 && x1 < m && y1 > 0 && y1 < n && grid[x1][y1] == '1' {
						v1 := x1*n + y1
						uf.Merge(v, v1)
					}
				}
			} else {
				waterCnt++
			}
		}
	}

	// 统计多少个集合，排除水格子
	return uf.NumOfSets() - waterCnt
}

func NumIslands_2(grid [][]byte) int {
	// 深度优先遍历
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// 从 i, j 开始探索，探索过的陆地标记为 2
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1' {
			grid[i][j] = '2'
			// 继续遍历上下左右
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 每次遇到 '1' 就说明遇见了新岛屿，然后将这个岛屿所有格子都探索一遍
			if grid[i][j] == '1' {
				dfs(i, j)
				cnt++
			}
		}
	}

	return cnt
}
