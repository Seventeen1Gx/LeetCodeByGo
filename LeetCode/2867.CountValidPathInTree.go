package LeetCode

func CountValidPathInTree_1(n int, edges [][]int) int64 {
	// 对于每个质数节点 p，考虑所有经过 p 的路径，且 p 是路径上唯一的质数：
	// - 路径可以以 p 为端点（p 本身是质数，另一端在非质数块中）
	// - 路径可以以 p 为中间点（两端都在不同的非质数块中）

	if n < 2 {
		return 0
	}

	// 先标记所有数都是质数
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	// 排除合数
	// 从某个质数开始，它的倍数都不是质数
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// 从 i*i 开始，因为小于这个数的 i 倍数之前都被处理过了
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 邻接表
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	// 计算非质数结点所在连通块的大小
	visited := make([]bool, n+1)
	size := make([]int, n+1)

	var dfs func(u, comID int) int
	dfs = func(u, comID int) int { // 从 u 开始，计算第 comID 个非质数连通块的大小
		visited[u] = true
		size[u] = comID // 记录 u 在哪个连通块中
		cnt := 1
		for _, v := range adj[u] {
			if !visited[v] && !isPrime[v] {
				cnt += dfs(v, comID)
			}
		}
		return cnt
	}

	// 给每个非质数连通块编号，并记录大小
	compSize := []int{0}
	comID := 0
	for i := 1; i <= n; i++ {
		if !isPrime[i] && !visited[i] {
			comID++
			sz := dfs(i, comID)
			compSize = append(compSize, sz)
		}
	}

	var ans int64 = 0

	// 枚举每个质数结点
	for i := 1; i <= n; i++ {
		if !isPrime[i] {
			continue
		}

		// 相邻非质数连通块大小
		comps := []int{}
		for _, v := range adj[i] {
			if !isPrime[v] {
				comps = append(comps, compSize[size[v]])
			}
		}

		// 统计路径
		// 情况1：以 i 为端点的路径
		sum := 0
		for _, c := range comps {
			sum += c
		}

		ans += int64(sum)

		// 情况2：路径经过 i，两端在不同的连通块
		prefixSum := 0 // 之前已经遍历到的连通块的总结点数
		for _, c := range comps {
			ans += int64(prefixSum * c)
			prefixSum += c
		}
	}

	return ans
}

func CountValidPathInTree_2(n int, edges [][]int) int64 {
	if n < 2 {
		return 0
	}

	// 先标记所有数都是质数
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0], isPrime[1] = false, false

	// 排除合数
	// 从某个质数开始，它的倍数都不是质数
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// 从 i*i 开始，因为小于这个数的 i 倍数之前都被处理过了
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 邻接表
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	// dp[i][0] 表示从 i 结点开始向下不包含质数的路径数（单个结点也算一个路径）
	// dp[i][1] 表示从 i 结点开始向下包含 1 个质数的路径数
	// 如果 i 不是质数，则 dp[i][0] = sum(dp[child][0])+1，dp[i][1] = sum(dp[child][1])
	// 如果 i 是质数，则 dp[i][0] = 0，dp[i][1] = sum(dp[child][0])+1

	dp0 := make([]int, n+1)
	dp1 := make([]int, n+1)
	var ans int64 = 0

	// 处理 dp 数组，当前处理到 u 结点
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		if isPrime[u] {
			dp0[u] = 0
			dp1[u] = 1
		} else {
			dp0[u] = 1
			dp1[u] = 0
		}

		// 以 u 为 LCA 的路径
		sum0, sum1 := 0, 0
		subPaths := []int{}

		for _, v := range adj[u] {
			if v == parent {
				// 避免回路
				continue
			}

			// 把儿子处理好，就能得到当前结点的值
			dfs(v, u)

			if isPrime[u] {
				dp1[u] += dp0[v]
			} else {
				dp0[u] += dp0[v]
				dp1[u] += dp1[v]
			}

			// 以 u 为 LCA 的路径
			if isPrime[u] {
				// 质数结点，收集子结点的dp0
				subPaths = append(subPaths, dp0[v])
			} else {
				// 非质数结点，统计一个子树有1质，一个子树0质
				ans += int64(sum0 * dp1[v]) // 之前0质 * 当前1质
				ans += int64(sum1 * dp0[v]) // 之前1质 * 当前0质
				sum0 += dp0[v]
				sum1 += dp1[v]
			}
		}

		// 处理质数结点
		if isPrime[u] {
			// 统计两端在不同子树
			sum := 0
			for _, cnt := range subPaths {
				ans += int64(sum * cnt)
				sum += cnt
			}
			// 加上u作为端点的路径
			ans += int64(sum)
		} else {
			// u作为端点到1质的子树路径数
			ans += int64(sum1)
		}
	}

	dfs(1, -1)
	return ans
}
