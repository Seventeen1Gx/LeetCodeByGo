package LeetCode

func FindAllAnagrams(s string, p string) []int {
	// 滑动窗口
	cnt1, cnt2 := [26]byte{}, [26]byte{}
	for _, ch := range p {
		cnt2[ch-'a']++
	}

	ans := []int{}
	left, right := 0, 0
	for right < len(s) {
		// 入
		cnt1[s[right]-'a']++

		// 还未形成第一个窗口
		if right-left+1 < len(p) {
			right++
			continue
		}

		// 更新
		if cnt1 == cnt2 {
			ans = append(ans, left)
		}

		// 出
		cnt1[s[left]-'a']--
		left++
		right++
	}

	return ans
}
