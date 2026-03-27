package LeetCode

import "LeetCodeByGo/utils"

func RhombusSum(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	diagSum := make([][]int, m+1) // ↘ 前缀和
	antiSum := make([][]int, m+1) // ↙ 前缀和
	for i := range diagSum {
		diagSum[i] = make([]int, n+1)
		antiSum[i] = make([]int, n+1)
	}
	for x, row := range grid {
		for y, v := range row {
			diagSum[x+1][y+1] = diagSum[x][y] + v
			// 在 grid 矩阵中 (x,y)  的位置，在 antiSum 矩阵中是 (x,y+1) 的位置，这个位置左下角是 (x+1,y)
			antiSum[x+1][y] = antiSum[x][y+1] + v
		}
	}

	utils.PrintMatrix(diagSum)
	utils.PrintMatrix(antiSum)

	// 从 grid 的 (x,y) 开始，向 ↘ 走 k 个元素的和，终点 (x+k-1,y+k-1)
	queryDiagSum := func(x, y, k int) int {
		return diagSum[x+k][y+k] - diagSum[x][y]
	}
	// 从 grid 的 (x,y) 开始，在 antiSum 矩阵中是 (x,y+1) 的位置，向 ↙ 走 k 个元素的和，终点 (x+k-1,y-k+2)，终点左下角 (x+k,y-k+1)
	queryAntiSum := func(x, y, k int) int {
		return antiSum[x+k][y-k+1] - antiSum[x][y+1]
	}

	// 相同的元素不更新
	var ans1, ans2, ans3 int // ans1 > ans2 > ans3
	updateAns := func(v int) {
		if v > ans1 {
			ans1, ans2, ans3 = v, ans1, ans2
		} else if v > ans2 && v < ans1 {
			ans2, ans3 = v, ans2
		} else if v < ans2 && v > ans3 {
			ans3 = v
		}
	}

	for x1 := 0; x1 < m; x1++ {
		for y1 := 0; y1 < n; y1++ {
			// 单个点的菱形
			updateAns(grid[x1][y1])

			// 以 (x,y) 为上顶点的其他菱形
			k := 2
			for {
				// 菱形需要在矩阵内
				x2, y2 := x1+k-1, y1-k+1
				x3, y3 := x1+k-1, y1+k-1
				x4, y4 := x2+k-1, y1
				if y2 < 0 || y3 >= n || x4 >= m {
					break
				}
				a := queryDiagSum(x1, y1, k)
				b := queryAntiSum(x1, y1, k)
				c := queryDiagSum(x2, y2, k)
				d := queryAntiSum(x3, y3, k)
				updateAns(a + b + c + d - grid[x1][y1] - grid[x2][y2] - grid[x3][y3] - grid[x4][y4])
				k++
			}
		}
	}

	var ans []int
	if ans1 > 0 {
		ans = append(ans, ans1)
	}
	if ans2 > 0 {
		ans = append(ans, ans2)
	}
	if ans3 > 0 {
		ans = append(ans, ans3)
	}
	return ans
}
