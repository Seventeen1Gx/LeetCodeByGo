package LeetCode

import "math"

func MinimumRecolor(blocks string, k int) int {
	n := len(blocks)
	ans := math.MaxInt
	whiteCnt := 0
	for right := 0; right < n; right++ {
		if blocks[right] == 'W' {
			whiteCnt++
		}
		left := right - k + 1
		if left < 0 {
			continue
		}
		ans = min(ans, whiteCnt)
		if blocks[left] == 'W' {
			whiteCnt--
		}
	}
	return ans
}
