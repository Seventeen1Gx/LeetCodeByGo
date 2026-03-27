package LeetCode

func RottingOranges(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])

	// 存储新鲜橘子
	hashSet := make(map[int]bool)
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				hashSet[i*n+j] = true
			}
		}
	}

	if len(hashSet) == 0 {
		return ans
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	type bfsNode struct {
		x, y, curMin int
	}

	queue := make([]*bfsNode, 0)

	for i, row := range grid {
		for j, v := range row {
			if v == 2 {
				queue = append(queue, &bfsNode{i, j, 0})
			}
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if len(hashSet) == 0 {
			ans = max(ans, node.curMin)
		}
		for _, dir := range dirs {
			x, y := node.x+dir[0], node.y+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
				grid[x][y] = 2
				delete(hashSet, x*n+y)
				queue = append(queue, &bfsNode{x, y, node.curMin + 1})
			}
		}
	}
	if len(hashSet) > 0 {
		return -1
	}

	return ans
}
