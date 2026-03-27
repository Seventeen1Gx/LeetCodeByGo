package LeetCode

func SplitMatrixEqualSum2(grid [][]int) bool {
	// 移除后，剩余部分要连通
	// 只有一行或一列的情况下，被移除在中间的格子，才会失去连通

	m, n := len(grid), len(grid[0])

	// hashSet[x] = {当前分块中 x 存在的列/行索引}
	hashSetB := make(map[int]map[[2]int]bool, m*n)

	sumB := 0
	for i, row := range grid {
		for j, v := range row {
			sumB += v
			if hashSetB[v] == nil {
				hashSetB[v] = map[[2]int]bool{{i, j}: true}
			} else {
				hashSetB[v][[2]int{i, j}] = true
			}
		}
	}

	// 按行分割
	sumA := 0
	hashSetA := make(map[int]map[[2]int]bool, m*n)
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			x := grid[i][j]
			sumA += x
			if hashSetA[x] == nil {
				hashSetA[x] = map[[2]int]bool{{i, j}: true}
			} else {
				hashSetA[x][[2]int{i, j}] = true
			}
			sumB -= x
			hashSetB[x][[2]int{i, j}] = false
		}
		if sumA == sumB {
			return true
		}
		if sumA > sumB {
			// 上方需要移除一个格子
			idxList := hashSetA[sumA-sumB]
			for idx, exist := range idxList {
				if !exist {
					continue
				}
				if i != 0 && n != 1 {
					// 上方部分至少是 2 * 2
					return true
				}
				if i == 0 && (idx[1] == 0 || idx[1] == n-1) {
					// 上方部分是 1*n 的情况
					return true
				}
				if n == 1 && (idx[0] == 0 || idx[0] == i) {
					// 上方部分是 n*1 的情况
					return true
				}
			}
		} else {
			// 下方需要移除一个格子
			idxList := hashSetB[sumB-sumA]
			for idx, exist := range idxList {
				if !exist {
					continue
				}
				if i != m-2 && n != 1 {
					// 下方部分至少是 2 * 2
					return true
				}
				if i == m-2 && (idx[1] == 0 || idx[1] == n-1) {
					// 下方部分是 1*n 的情况
					return true
				}
				if n == 1 && (idx[0] == i+1 || idx[0] == m-1) {
					// 下方部分是 n*1 的情况
					return true
				}
			}
		}
	}

	// 按列分割
	hashSetB = make(map[int]map[[2]int]bool, m*n)

	sumB = 0
	for i, row := range grid {
		for j, v := range row {
			sumB += v
			if hashSetB[v] == nil {
				hashSetB[v] = map[[2]int]bool{{i, j}: true}
			} else {
				hashSetB[v][[2]int{i, j}] = true
			}
		}
	}

	sumA = 0
	hashSetA = make(map[int]map[[2]int]bool, m*n)
	for j := 0; j < n-1; j++ {
		for i := 0; i < m; i++ {
			x := grid[i][j]
			sumA += x
			if hashSetA[x] == nil {
				hashSetA[x] = map[[2]int]bool{{i, j}: true}
			} else {
				hashSetA[x][[2]int{i, j}] = true
			}
			sumB -= x
			hashSetB[x][[2]int{i, j}] = false
		}
		if sumA == sumB {
			return true
		}
		if sumA > sumB {
			// 左边需要移除一个格子
			idxList := hashSetA[sumA-sumB]
			for idx, exist := range idxList {
				if !exist {
					continue
				}
				if j != 0 && m != 1 {
					return true
				}
				if j == 0 && (idx[0] == 0 || idx[0] == m-1) {
					return true
				}
				if m == 1 && (idx[1] == 0 || idx[1] == j) {
					return true
				}
			}
		} else {
			// 右边需要移除一个格子
			idxList := hashSetB[sumB-sumA]
			for idx, exist := range idxList {
				if !exist {
					continue
				}
				if j != n-2 && m != 1 {
					return true
				}
				if j == m-2 && (idx[0] == 0 || idx[0] == m-1) {
					return true
				}
				if m == 1 && (idx[1] == j+1 || idx[1] == n-1) {
					return true
				}
			}
		}
	}

	return false
}
