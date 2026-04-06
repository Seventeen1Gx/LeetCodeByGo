package LeetCode

import "strconv"

func CountAndSay(n int) string {
	RLE := func(s string) string {
		n := len(s)
		if n == 0 {
			return ""
		}

		ans := ""
		cnt := 0
		cntCh := ""
		for i := 0; i < n; i++ {
			if i == 0 || s[i-1] != s[i] {
				// 首次遇见某个字符
				if cnt != 0 {
					// 处理上一个字符
					ans += strconv.Itoa(cnt) + cntCh
				}
				cnt = 1
				cntCh = s[i : i+1]
			} else {
				// 重复遇见字符
				cnt++
			}
		}
		ans += strconv.Itoa(cnt) + cntCh
		return ans
	}

	pre := "1"
	cur := ""
	for i := 1; i < n; i++ {
		cur = RLE(pre)
		pre = cur
	}

	return pre
}
