package LeetCode

import (
	"sort"
)

func CheckEdgeLimitPathExist(n int, edgeList [][]int, queries [][]int) []bool {
	// 判断两点是否连通 -> 并查集
	us := make([]int, n)
	for i := range us {
		us[i] = -1
	}

	var find func(x int) int
	find = func(x int) int {
		if us[x] < 0 {
			return x
		}
		y := find(us[x])
		us[x] = y
		return y
	}

	merge := func(x, y int) {
		rootx := find(x)
		rooty := find(y)
		if rootx != rooty {
			us[rootx] = rooty
		}
	}

	canReach := func(from, to int) bool {
		return find(from) == find(to)
	}

	// 为了排序后还能找到原来的索引值
	m := len(queries)
	id := make([]int, m)
	for i := range id {
		id[i] = i
	}

	// 从小到大取边 edgeList
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	// 从小到达处理 queries
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]][2] < queries[id[j]][2]
	})

	edgeIdx := 0
	ans := make([]bool, m)
	for _, i := range id {
		query := queries[i]
		p, q, l := query[0], query[1], query[2]
		// 将小于 l 的边加入并查集
		for edgeIdx < len(edgeList) {
			edge := edgeList[edgeIdx]
			u, v, dist := edge[0], edge[1], edge[2]
			if dist >= l {
				break
			}
			merge(u, v)
			edgeIdx++
		}
		// 判断 p, q 是否连通
		ans[i] = canReach(p, q)
	}

	return ans
}
