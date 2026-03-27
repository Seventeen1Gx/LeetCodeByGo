package LeetCode

func SlipSquareMatrix(grid [][]int, x int, y, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		for j := y; j < y+k; j++ {
			grid[i][j], grid[2*x+k-i-1][j] = grid[2*x+k-i-1][j], grid[i][j]

		}
	}
	return grid
}
