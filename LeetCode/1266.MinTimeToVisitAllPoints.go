package LeetCode

import (
	"LeetCodeByGo/utils"
)

func MinTimeToVisitAllPoints(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	var time int
	for i := 0; i < n-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[i+1][0], points[i+1][1]
		dx := utils.Abs(x1 - x2)
		dy := utils.Abs(y1 - y2)
		time += max(dx, dy)
	}

	return time
}
