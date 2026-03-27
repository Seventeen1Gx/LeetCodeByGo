package LeetCode

import "container/heap"

// 二维情况下，并不是 4 个方向上最大值就行了
// 而是一圈一圈考虑，每圈考虑最低点
func TrapRain2(heightMap [][]int) int {
	ans := 0
	m, n := len(heightMap), len(heightMap[0])

	hp := cellHeap{}
	// 将最外一圈加入堆中
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(&hp, &cell{i, j, v})
				row[j] = -1 // 标记已访问
			}
		}
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for hp.Len() > 0 {
		c := heap.Pop(&hp).(*cell) // 当前最短板
		minHeight, i, j := c.height, c.i, c.j
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && heightMap[x][y] >= 0 {
				// 未访问过的邻居
				ans += max(minHeight-heightMap[x][y], 0)
				// 短板出去了，需要补充根
				heap.Push(&hp, &cell{x, y, max(minHeight, heightMap[x][y])})
				// 标记访问过
				heightMap[x][y] = -1
			}
		}
	}

	return ans
}

type cell struct {
	i, j, height int
}

type cellHeap struct {
	cells []*cell
}

func (c *cellHeap) Len() int           { return len(c.cells) }
func (c *cellHeap) Less(i, j int) bool { return c.cells[i].height < c.cells[j].height }
func (c *cellHeap) Swap(i, j int)      { c.cells[i], c.cells[j] = c.cells[j], c.cells[i] }
func (c *cellHeap) Push(x interface{}) { c.cells = append(c.cells, x.(*cell)) }
func (c *cellHeap) Pop() interface{} {
	n := c.Len()
	x := c.cells[n-1]
	c.cells = c.cells[:n-1]
	return x
}
