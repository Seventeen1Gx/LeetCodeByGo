package LeetCode

import "strings"

func MaxVowels(s string, k int) int {
	n := len(s)
	left := 0
	right := 0
	window := 0
	maxV := 0
	for right < k {
		if strings.Contains("aeiou", s[right:right+1]) {
			window++
			maxV = max(window, maxV)
		}
		right++
	}
	for right < n {
		if strings.Contains("aeiou", s[right:right+1]) {
			window++
		}
		if strings.Contains("aeiou", s[left:left+1]) {
			window--
		}
		maxV = max(window, maxV)
		left++
		right++
	}
	return maxV
}
