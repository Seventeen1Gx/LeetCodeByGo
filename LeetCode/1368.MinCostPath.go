package LeetCode

import "container/heap"

func MinCost2MakeAtLeastValidPath_1(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// 单源最短路径问题
	// m * n 个结点
	// 每个结点有 4 条边连接到上下左右的结点上
	// 如果边方向与结点本身相同，则这条边的权为 0 否则为 1

	dist := make([]int, m*n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0
	visited := make([]bool, m*n)

	queue := &NodeHeap{}
	heap.Push(queue, &NodeDist{0, 0})
	for queue.Len() > 0 {
		node := heap.Pop(queue).(*NodeDist)
		if visited[node.idx] {
			continue
		}
		visited[node.idx] = true

		x, y := node.idx/n, node.idx%n
		for i, dir := range dirs {
			weight := 1
			if i+1 == grid[x][y] {
				weight = 0
			}
			if grid[x][y] > 4 || grid[x][y] < 1 {
				continue
			}

			xx, yy := x+dir[0], y+dir[1]
			if xx < 0 || xx >= m || yy < 0 || yy >= n {
				continue
			}

			newDist := dist[node.idx] + weight
			if dist[xx*n+yy] < 0 || dist[xx*n+yy] > newDist {
				dist[xx*n+yy] = newDist
				heap.Push(queue, &NodeDist{xx*n + yy, newDist})
			}
		}
	}

	return dist[len(dist)-1]
}
