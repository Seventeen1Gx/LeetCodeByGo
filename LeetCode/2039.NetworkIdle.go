package LeetCode

func NetworkIdle(edges [][]int, patience []int) int {
	// 变为空闲状态的最早秒数:每条消息都得到了回复（只有第一条消息没回复，才会继续重传）

	// 先计算主服务器到其他服务器的路径长度 -> BFS
	n := len(patience)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	d := 0
	queue := make([]int, 0)
	dist := make([]int, n)
	visited := make([]bool, n)
	queue = append(queue, 0)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			u := queue[0]
			queue = queue[1:]
			if visited[u] {
				continue
			}
			dist[u] = d
			visited[u] = true

			for _, v := range adj[u] {
				if !visited[v] {
					queue = append(queue, v)
				}
			}
		}
		d++
	}

	ans := 0
	for i, p := range patience {
		if i == 0 {
			continue
		}
		ttl := dist[i] * 2
		if p >= ttl {
			// 不会重发消息
			ans = max(ans, ttl+1)
		} else {
			// 每隔 p 秒重发一次消息
			// 在第一次消息返回之前，最后一次重发消息走完后网络空闲
			// ttl=9 p=3 在第 9 秒时不会重发，最后一次是第 6 秒
			ans = max(ans, (ttl-1)/p*p+ttl+1)
		}
	}

	return ans
}
