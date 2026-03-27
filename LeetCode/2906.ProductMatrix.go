package LeetCode

func ProductMatrix_1(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])

	p := make([][]int, m)
	for i := range p {
		p[i] = make([]int, n)
	}

	total := 1
	for _, row := range grid {
		for _, v := range row {
			total *= v
		}
	}

	for i, row := range grid {
		for j, v := range row {
			p[i][j] = (total / v) % 12345
		}
	}

	return p
}

func ProductMatrix_2(grid [][]int) [][]int {
	// 参考 238 题
	// 二维变一维
	m, n := len(grid), len(grid[0])

	pre, suf := make([]int, m*n), make([]int, m*n)
	pre[0], suf[m*n-1] = 1, 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			if idx == 0 {
				continue
			}
			ii, jj := (idx-1)/n, (idx-1)%n
			pre[idx] = pre[idx-1] * grid[ii][jj] % 12345
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			idx := i*n + j
			if idx == m*n-1 {
				continue
			}
			ii, jj := (idx+1)/n, (idx+1)%n
			suf[idx] = suf[idx+1] * grid[ii][jj] % 12345
		}
	}

	p := make([][]int, m)
	for i := range p {
		p[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			p[i][j] = pre[idx] * suf[idx] % 12345
		}
	}

	return p
}
