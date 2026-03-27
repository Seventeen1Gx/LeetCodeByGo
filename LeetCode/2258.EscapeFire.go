package LeetCode

import "sort"

func EscapeFire_1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	type bfsNode struct {
		x, y int
	}

	canReach := func(grid [][]int, queueFire []*bfsNode) bool {
		// 能否从矩阵左上角到达右下角
		if grid[0][0] != 0 || grid[m-1][n-1] != 0 {
			return false
		}

		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		// 入队列说明到达那个点
		queuePeople := make([]*bfsNode, 0)
		queuePeople = append(queuePeople, &bfsNode{0, 0})
		visited[0][0] = true

		for len(queuePeople) > 0 {
			// 先走再烧，但是如果出队列的时候发现当前所站的点被烧了就跳过
			size := len(queuePeople)
			for i := 0; i < size; i++ {
				node := queuePeople[0]
				queuePeople = queuePeople[1:]
				if grid[node.x][node.y] == 1 {
					continue
				}
				for _, dir := range dirs {
					x, y := node.x+dir[0], node.y+dir[1]
					if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && grid[x][y] == 0 {
						visited[x][y] = true
						if x == m-1 && y == n-1 {
							return true
						}
						queuePeople = append(queuePeople, &bfsNode{x, y})
					}
				}
			}

			size = len(queueFire)
			for i := 0; i < size; i++ {
				nodeFire := queueFire[0]
				queueFire = queueFire[1:]
				for _, dir := range dirs {
					x, y := nodeFire.x+dir[0], nodeFire.y+dir[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 {
						grid[x][y] = 1
						queueFire = append(queueFire, &bfsNode{x, y})
					}
				}
			}
		}

		return false
	}

	// 火势蔓延队列
	// 入队列说明到达那个点
	queueFire := make([]*bfsNode, 0)
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			queueFire = append(queueFire, &bfsNode{i, j})
		}
	}
	if len(queueFire) == 0 {
		if canReach(grid, nil) {
			return 1_000_000_000
		} else {
			return -1
		}
	}

	curMin := 0
	for len(queueFire) > 0 {
		gridCopy := make([][]int, m)
		for i := range gridCopy {
			gridCopy[i] = make([]int, n)
			copy(gridCopy[i], grid[i])
		}
		if !canReach(gridCopy, queueFire) {
			return curMin - 1
		}
		size := len(queueFire)
		for i := 0; i < size; i++ {
			node := queueFire[0]
			queueFire = queueFire[1:]
			for _, dir := range dirs {
				x, y := node.x+dir[0], node.y+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 {
					grid[x][y] = 1
					queueFire = append(queueFire, &bfsNode{x, y})
				}
			}
		}
		curMin++
		// 每过一分钟，检查从此刻开始逃脱是否可以到达终点
	}

	return 1_000_000_000
}

func EscapeFire_2(grid [][]int) int {
	// 二分查找
	m, n := len(grid), len(grid[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	type bfsNode struct {
		x, y int
	}

	var queueFire []*bfsNode

	fireMin := func(grid [][]int) {
		// 燃烧一分钟
		size := len(queueFire)
		for i := 0; i < size; i++ {
			nodeFire := queueFire[0]
			queueFire = queueFire[1:]
			for _, dir := range dirs {
				x, y := nodeFire.x+dir[0], nodeFire.y+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 0 {
					grid[x][y] = 1
					queueFire = append(queueFire, &bfsNode{x, y})
				}
			}
		}
	}

	canReach := func(grid [][]int, t int) bool {
		// 初始火焰
		queueFire = make([]*bfsNode, 0)
		for i, row := range grid {
			for j, v := range row {
				if v != 1 {
					continue
				}
				queueFire = append(queueFire, &bfsNode{i, j})
			}
		}

		// 经过 t 分钟的火势蔓延后，能否从矩阵左上角到达右下角
		for len(queueFire) > 0 && t > 0 {
			fireMin(grid)
			if grid[0][0] != 0 || grid[m-1][n-1] != 0 {
				return false
			}
			t--
		}

		if grid[0][0] != 0 || grid[m-1][n-1] != 0 {
			return false
		}

		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		// 入队列说明到达那个点
		queuePeople := make([]*bfsNode, 0)
		queuePeople = append(queuePeople, &bfsNode{0, 0})
		visited[0][0] = true

		for len(queuePeople) > 0 {
			// 先走再烧，但是如果出队列的时候发现当前所站的点被烧了就跳过
			size := len(queuePeople)
			for i := 0; i < size; i++ {
				node := queuePeople[0]
				queuePeople = queuePeople[1:]
				if grid[node.x][node.y] == 1 {
					continue
				}
				for _, dir := range dirs {
					x, y := node.x+dir[0], node.y+dir[1]
					if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && grid[x][y] == 0 {
						visited[x][y] = true
						if x == m-1 && y == n-1 {
							return true
						}
						queuePeople = append(queuePeople, &bfsNode{x, y})
					}
				}
			}

			fireMin(grid)
		}

		return false
	}

	// 二分寻找满足条件的最大 t 值
	// Search 函数返回满足 f(i) = true 的最小索引 i
	// 这里得到的 t 就是不能到达终点的最小的 t，再小就可以到达终点
	t := sort.Search(m*n+1, func(t int) bool {
		gridCopy := make([][]int, m)
		for i := range gridCopy {
			gridCopy[i] = make([]int, n)
			copy(gridCopy[i], grid[i])
		}
		return !canReach(gridCopy, t)
	})

	if t == m*n+1 {
		return 1_000_000_000
	}
	if t == 0 {
		return -1
	}
	return t - 1
}
