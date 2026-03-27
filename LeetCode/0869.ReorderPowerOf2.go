package LeetCode

import (
	"slices"
	"strconv"
)

func ReorderPowerOf2(n int) bool {
	// 预处理 2 的幂
	anagramsSet := make(map[string]bool)
	for i := 1; i < 1_000_000_000; i <<= 1 {
		strI := strconv.Itoa(i)
		tmpI := []byte(strI)
		slices.Sort(tmpI)
		anagramsSet[string(tmpI)] = true
	}

	// 看给的数字是否与 2 的幂是异位词
	strN := strconv.Itoa(n)
	tmpN := []byte(strN)
	slices.Sort(tmpN)
	return anagramsSet[string(tmpN)]
}
