package LeetCode

import "strings"

func MaxNumVowelsInSubstring(s string, k int) int {
	cnt, maxCnt := 0, 0
	for right := 0; right < len(s); right++ {
		if strings.Contains("aeiou", s[right:right+1]) {
			cnt++
		}

		left := right + 1 - k
		if left < 0 {
			continue
		}

		if cnt > maxCnt {
			maxCnt = cnt
		}

		if strings.Contains("aeiou", s[left:left+1]) {
			cnt--
		}
	}

	return maxCnt
}
