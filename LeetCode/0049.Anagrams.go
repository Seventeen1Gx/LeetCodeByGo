package LeetCode

import (
	"slices"
)

func GroupAnagrams(strs []string) [][]string {
	// 排序字符串
	var hashSet = make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		slices.Sort(tmp)
		sortS := string(tmp)
		hashSet[sortS] = append(hashSet[sortS], str)
	}

	var ans = make([][]string, 0, len(hashSet))
	for _, v := range hashSet {
		ans = append(ans, v)
	}
	return ans
}
