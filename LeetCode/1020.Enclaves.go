package LeetCode

import "LeetCodeByGo/utils"

func Enclaves_1(grid [][]int) int {
	// DFS:从边界陆地出发标记与之相连的陆地
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		grid[i][j] = -1 // 标记可到达
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 1 {
			dfs(i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			dfs(0, j)
		}
		if grid[m-1][j] == 1 {
			dfs(m-1, j)
		}
	}

	ans := 0
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				ans++
			}
		}
	}
	return ans
}

func Enclaves_2(grid [][]int) int {
	// BFS:从边界陆地出发标记与之相连的陆地
	// 注意这里要入队的时候标记可到达，不然在队列中等待的时候，被其他结点再进入了
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	queue := make([][]int, 0)
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			queue = append(queue, []int{i, 0})
		}
		if grid[i][n-1] == 1 {
			queue = append(queue, []int{i, n - 1})
		}
	}

	for j := 0; j < n; j++ {
		if grid[0][j] == 1 {
			queue = append(queue, []int{0, j})
		}
		if grid[m-1][j] == 1 {
			queue = append(queue, []int{m - 1, j})
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		i, j := node[0], node[1]
		if grid[i][j] != 1 {
			// 结点等待过程中可能重复入队
			continue
		}
		grid[i][j] = -1 // 标记可到达
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
				queue = append(queue, []int{x, y})
			}
		}
	}

	ans := 0
	for _, row := range grid {
		for _, v := range row {
			if v == 1 {
				ans++
			}
		}
	}
	return ans
}

func Enclaves_3(grid [][]int) int {
	// 并查集
	m, n := len(grid), len(grid[0])

	top, bottom, left, right := m*n, m*n+1, m*n+2, m*n+3
	uf := utils.NewUnionFind(m*n + 4)

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			visited[i][j] = true
			if i == 0 {
				uf.Merge(i*n+j, top)
			}
			if i == m-1 {
				uf.Merge(i*n+j, bottom)
			}
			if j == 0 {
				uf.Merge(i*n+j, left)
			}
			if j == n-1 {
				uf.Merge(i*n+j, right)
			}
			for _, dir := range dirs {
				x, y := i+dir[0], j+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 && !visited[x][y] {
					uf.Merge(i*n+j, x*n+y)
				}
			}
		}
	}

	ans := 0
	rootA, rootB, rootC, rootD := uf.Find(top), uf.Find(bottom), uf.Find(left), uf.Find(right)
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				continue
			}
			rootE := uf.Find(i*n + j)
			if rootE != rootA && rootE != rootB && rootE != rootC && rootE != rootD {
				ans++
			}
		}
	}

	return ans
}
