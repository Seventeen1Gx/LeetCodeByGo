package LeetCode

func MinimizeTotalPriceOfTrip(n int, edges [][]int, price []int, trips [][]int) int {
	// 因为是树，两点之间只有唯一路径
	// 先统计整个旅行下来，每个结点被访问多少次

	// 邻接表
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	getPath := func(s, t int) []int {
		from := make([]int, n)
		for i := range from {
			from[i] = -1
		}
		from[s] = s
		// from[i] = j 说明从 j 到 i

		queue := make([]int, 0)
		queue = append(queue, s)
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			if u == t {
				break
			}
			for _, v := range adj[u] {
				if from[v] == -1 {
					from[v] = u
					queue = append(queue, v)
				}
			}
		}

		path := make([]int, 0)
		for u := t; ; u = from[u] {
			path = append(path, u)
			if u == s {
				break
			}
		}
		return path
	}

	cnt := make([]int, n)
	for _, trip := range trips {
		path := getPath(trip[0], trip[1])
		for _, u := range path {
			cnt[u]++
		}
	}

	// 每个结点的权重
	w := make([]int, n)
	for i := range w {
		w[i] = price[i] * cnt[i]
	}

	// 现在要选择一些不相邻的点将其 w 减半
	// 求最小总花费

	// dp[u][0] 表示结点 u 不减半，以 u 为根的子树的最小花费
	// dp[u][1] 表示结点 u 减半，以 u 为根的子树的最小花费
	// dp[u][0] = w[u] + sum(min(dp[v][0], dp[v][1]))
	// dp[u][1] = w[u]/2 + sum(dp[v][0])
	var dfs func(u, parent int) (int, int)
	dfs = func(u, parent int) (int, int) {
		notSelect := w[u]     // 结点 u 不减半
		selectVal := w[u] / 2 // 结点 u 减半
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			ns, s := dfs(v, u)
			notSelect += min(ns, s)
			selectVal += ns
		}
		return notSelect, selectVal
	}

	ns, s := dfs(0, -1)
	return min(ns, s)
}
