package LeetCode

func MaxRectangle(matrix [][]byte) (ans int) {
	// 枚举矩形的底，套用上一题的做法
	_, m := len(matrix), len(matrix[0])
	heights := make([]int, m)
	for _, row := range matrix {
		for i, v := range row {
			if v == '1' {
				heights[i] += 1
			} else {
				heights[i] = 0
			}
		}
		ans = max(ans, MaxRectangleInHistogram(heights))
	}

	return ans
}
