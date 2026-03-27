package LeetCode

import "sort"

func MatrixMaxScore(grid [][]int, queries []int) []int {
	// 矩阵每个格子是图的一个结点，值严格小于 queries[i] 的结点与上下左右同样小于 queries[i] 的结点相连
	// 找到那些与左上角结点相连的结点
	m, n, k := len(grid), len(grid[0]), len(queries)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	uf := make([]int, m*n)
	for i := range uf {
		uf[i] = -1
	}

	var find func(x int) int
	find = func(x int) int {
		if uf[x] < 0 {
			return x
		}
		y := find(uf[x])
		uf[x] = y
		return y
	}

	merge := func(x, y int) {
		rootx := find(x)
		rooty := find(y)
		if rootx != rooty {
			sizex := -uf[rootx]
			sizey := -uf[rooty]
			if sizex < sizey {
				uf[rootx] = rooty
				uf[rooty] -= sizex
			} else {
				uf[rooty] = rootx
				uf[rootx] -= sizey
			}
		}
	}

	size := func(x int) int {
		rootx := find(x)
		return -uf[rootx]
	}

	node := make([][3]int, 0, m*n)
	for i, row := range grid {
		for j, v := range row {
			node = append(node, [3]int{i, j, v})
		}
	}

	sort.Slice(node, func(i, j int) bool {
		return node[i][2] < node[j][2]
	})

	id := make([]int, k)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]] < queries[id[j]]
	})

	j := 0
	answer := make([]int, k)
	for _, i := range id {
		query := queries[i]
		// 小于 query 的结点加入并查集
		for j < m*n {
			if node[j][2] >= query {
				break
			}
			x1, y1 := node[j][0], node[j][1]
			for _, dir := range dirs {
				x2, y2 := x1+dir[0], y1+dir[1]
				if x2 >= 0 && x2 < m && y2 >= 0 && y2 < n && grid[x2][y2] < query {
					merge(x1*n+y1, x2*n+y2)
				}
			}
			j++
		}
		if grid[0][0] < query {
			answer[i] = size(0)
		}
	}

	return answer
}
