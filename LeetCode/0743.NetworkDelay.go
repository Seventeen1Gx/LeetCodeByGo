package LeetCode

import (
	"container/heap"
	"math"
)

func NetworkDelay(times [][]int, n int, k int) int {
	// 迪杰斯特拉算法，在图中从某个结点出发，到达其他每个结点的最短路径
	// 要求所有边权重非负
	// 每次选择距离源点最近的“未处理”节点，确认它的最短距离，然后用它去松弛（更新）相邻节点的距离。

	// 构造邻接点表
	adjList := make(map[int][]*NodeDist)
	for _, time := range times {
		adjList[time[0]-1] = append(adjList[time[0]-1], &NodeDist{
			idx:  time[1] - 1,
			dist: time[2],
		})
	}

	// dist[v] 表示当前从源点到 v 点的最短距离
	// 初始化时，dist[源点]=0，dist[其他点]=∞，因为此时其他点都不能到达

	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k-1] = 0

	// visited[v] 表示结点 v 的最短距离是否已经确认
	visited := make([]bool, n)

	hp := NodeHeap{}
	heap.Push(&hp, &NodeDist{idx: k - 1, dist: 0})

	for hp.Len() > 0 {
		// 在未访问的结点中选择 dest 值最小的
		item := heap.Pop(&hp).(*NodeDist)

		if visited[item.idx] {
			continue
		}
		visited[item.idx] = true

		// 对于它的邻接点
		for _, node := range adjList[item.idx] {
			if dist[node.idx] > dist[item.idx]+node.dist {
				dist[node.idx] = dist[item.idx] + node.dist
				heap.Push(&hp, &NodeDist{idx: node.idx, dist: dist[item.idx] + node.dist})
			}
		}
	}

	mx := 0
	for _, v := range dist {
		mx = max(mx, v)
	}

	if mx == math.MaxInt {
		return -1
	}
	return mx
}

type NodeDist struct {
	idx  int
	dist int
}

type NodeHeap struct {
	dests []*NodeDist
}

func (hp *NodeHeap) Len() int           { return len(hp.dests) }
func (hp *NodeHeap) Less(i, j int) bool { return hp.dests[i].dist < hp.dests[j].dist }
func (hp *NodeHeap) Swap(i, j int)      { hp.dests[i], hp.dests[j] = hp.dests[j], hp.dests[i] }
func (hp *NodeHeap) Push(x interface{}) { hp.dests = append(hp.dests, x.(*NodeDist)) }
func (hp *NodeHeap) Pop() interface{} {
	n := len(hp.dests)
	x := hp.dests[n-1]
	hp.dests = hp.dests[:n-1]
	return x
}
