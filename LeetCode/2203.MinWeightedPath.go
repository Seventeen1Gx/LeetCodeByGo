package LeetCode

import (
	"container/heap"
)

func MinWeightedPath(n int, edges [][]int, src1 int, src2 int, dest int) int64 {
	// 最终的结果要么 src1 -> .. -> dest <- ... <- src2 没有公共边
	// 要么 src1 -> ... -> x -> ... -> dest	在 x 相遇后到达终点
	// 		src2 -> ... -> x
	// 它们不会相遇后离开，因为如果相遇后离开，删除相遇后离开再到达终点的一条路径，仍然可以到达终点
	// 删除后的图肯定比原先的图小

	// 我们枚举相遇点，再计算 src1、src2、dest 到相遇点的最短距离
	// 求三个最短距离和的最小值

	// 邻接表
	adj := make(map[int][][]int)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := edge[2]
		adj[from] = append(adj[from], []int{to, weight})
	}
	// 反向邻接表
	revAdj := make(map[int][][]int)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := edge[2]
		revAdj[to] = append(revAdj[to], []int{from, weight})
	}

	minPath := func(from int, adj map[int][][]int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[from] = 0

		visited := make([]bool, n)

		hp := NodeHeap{}
		heap.Push(&hp, &NodeDist{idx: from, dist: 0})
		for hp.Len() > 0 {
			nodeA := heap.Pop(&hp).(*NodeDist).idx
			if visited[nodeA] {
				continue
			}
			visited[nodeA] = true

			for _, edge := range adj[nodeA] {
				nodeB := edge[0]
				newDist := dist[nodeA] + edge[1]
				if dist[nodeB] == -1 || newDist < dist[nodeB] {
					dist[nodeB] = newDist
					heap.Push(&hp, &NodeDist{idx: nodeB, dist: newDist})
				}
			}
		}

		return dist
	}

	distA := minPath(src1, adj)
	distB := minPath(src2, adj)
	distC := minPath(dest, revAdj)

	ans := -1
	for i := 0; i < n; i++ {
		a := distA[i]
		if a < 0 {
			continue
		}
		b := distB[i]
		if b < 0 {
			continue
		}
		c := distC[i]
		if c < 0 {
			continue // 剪枝优化
		}
		if ans < 0 || a+b+c < ans {
			ans = a + b + c
		}
	}

	return int64(ans)
}
