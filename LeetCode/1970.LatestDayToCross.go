package LeetCode

// 并查集：高效处理集合的合并与查找
// 1. 用树表示集合，树的根节点就是集合的代表元
// 2. 主要两个操作：Find 与 Merge
// 3. 刚开始用一维数组表示树，fa[i]=j 表示元素 j 是元素 i 的父亲，但经过路径压缩，fa[i]=j 表示元素 i 所在集合的代表元是 j

type unionFind struct {
	fa []int // 代表元
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)

	// 一开始有 n 个集合 {0},{1}, ... ,{n-1}
	// 集合 i 的代表元是自己
	for i := range fa {
		fa[i] = i
	}

	return unionFind{fa}
}

// 返回元素 x 所在的元
// 同时作路径压缩，也就是把 x 所在集合的所有元素的 fa 都改成代表元
func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	// 如果 fa[x] == x，则表示 x 是代表元
	return u.fa[x]
}

// 判断 x 和 y 是否在同一集合
func (u unionFind) same(x, y int) bool {
	// 判断 x 和 y 的代表元是否相同（代表元表示对应集合）
	return u.find(x) == u.find(y)
}

// 将 from 集合合并到 to 中
func (u *unionFind) merge(from, to int) {
	x, y := u.find(from), u.find(to) // 找到代表元
	// 将一个树的根节点指向另一棵树的根节点
	u.fa[x] = y
}

func latestDayToCross1(row int, col int, cells [][]int) int {
	// 增加两个结点，表示上边的上面和下边的下面
	top := row * col
	bottom := row*col + 1
	uf := newUnionFind(row*col + 2)

	// land[i][j] 是否为陆地
	land := make([][]bool, row)
	for i := range land {
		land[i] = make([]bool, col)
	}

	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	// 从最后一天往前推
	for day := len(cells) - 1; day >= 0; day-- {
		cell := cells[day]
		r, c := cell[0]-1, cell[1]-1
		land[r][c] = true // 变成陆地

		// 并查集结点编号
		v := r*col + c

		if r == 0 {
			// 与上面的结点相连
			uf.merge(v, top)
		}

		if r == row-1 {
			// 与下面的结点相连
			uf.merge(v, bottom)
		}

		// 与上下左右的陆地结点相连
		for _, d := range dirs {
			rr, cc := r+d.x, c+d.y
			vv := rr*col + cc
			if rr >= 0 && rr < row && cc >= 0 && cc < col && land[rr][cc] {
				uf.merge(v, vv)
			}
		}

		if uf.same(top, bottom) {
			// 上下相连
			return day
		}
	}

	return -1
}

func latestDayToCross2(row int, col int, cells [][]int) int {
	// state = 0 ：全是水的格子
	// state = 1 ：刚变成陆地的格子
	// state = 2 : 确认能从顶部连通的格子
	state := make([][]int8, row)
	for i := range state {
		state[i] = make([]int8, col)
	}

	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	// 能否从第一行到达 (r, c)
	canReachFromTop := func(r, c int) bool {
		if r == 0 {
			// 已经是第一行
			return true
		}
		for _, d := range dirs {
			// 上下左右存在能从顶部到达的格子
			rr, cc := r+d.x, c+d.y
			if rr >= 0 && rr < row && cc >= 0 && cc < col && state[rr][cc] == 2 {
				return true
			}
		}
		return false
	}

	// 已经确认从顶部能到达 (r, c)，那么从 (r, c) 出发，能否到达最后一行
	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r == row-1 {
			return true
		}
		// 标记自己被“感染”
		state[r][c] = 2
		// “感染”上下左右的陆地格子
		for _, d := range dirs {
			rr, cc := r+d.x, c+d.y
			if rr >= 0 && rr < row && cc >= 0 && cc < col && state[rr][cc] == 1 && dfs(rr, cc) {
				return true
			}
		}
		return false
	}

	for day := len(cells) - 1; day >= 0; day-- {
		cell := cells[day]
		r, c := cell[0]-1, cell[1]-1 // 下标从零开始
		state[r][c] = 1
		if canReachFromTop(r, c) && dfs(r, c) {
			return day
		}
	}

	return -1
}
