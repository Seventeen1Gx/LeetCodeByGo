package LeetCode

// 对 n 行 3 列矩阵进行涂色，有三种颜色，保证相邻块不同色

const mod = 1_000_000_007

func GridColoring_1(n int) int {
	// 暴力深搜
	var grid = make([][]uint8, n)
	for i := range grid {
		grid[i] = make([]uint8, 3)
	}

	var ans int

	// 从上往下，从左往右进行涂色
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x == n {
			// 涂完所有颜色
			ans++
			ans %= mod
			return
		}

		// 对 i,j 进行涂色，保证跟他的左边和上面不一样
		for color := uint8(1); color <= 3; color++ {
			if (x == 0 || grid[x-1][y] != color) && (y == 0 || grid[x][y-1] != color) {
				grid[x][y] = color
				// 涂下一块
				if y == 2 {
					dfs(x+1, 0)
				} else {
					dfs(x, y+1)
				}
			}
		}
	}

	dfs(0, 0)
	return ans
}

func GridColoring_2(n int) int {
	// 递推解法
	// 一行要么是 'ABA' 类型，要么是 'ABC' 类型
	// 'ABA' : 3 * 2 = 6 种
	// 'ABC' : 3 * 2 * 1 = 6 种
	// 当上一行是 'ABA' 类型，下一行有 'ABA' 型的 3 种 : ABA -> BAB BCB CAC，下一行有 'ABC' 型的 2 种 : ABA -> BAC CAB
	// 当上一行是 'ABC' 类型，下一行有 'ABA' 型的 2 种 : ABC -> BAB BCB，下一行有 'ABC' 型的 2 种 : ABC -> BCA CAB

	// ABA 表示以 ABA 类型结尾的涂色方式，ABC 表示以 ABC 类型结尾的涂色方式

	// n = 1 时
	ABA := 6
	ABC := 6
	// 从 n=2 开始，迭代获得对应 n 的涂色方式
	for i := 2; i <= n; i++ {
		newABA := (ABA*3 + ABC*2) % mod
		newABC := (ABA*2 + ABC*2) % mod
		ABA = newABA
		ABC = newABC
	}

	return (ABA + ABC) % mod
}
