package LeetCode

func PermutationInString(s1 string, s2 string) bool {
	cnt1, cnt2 := [26]byte{}, [26]byte{}

	for _, v := range s2 {
		cnt2[v-'a']++
	}

	for right := 0; right < len(s1); right++ {
		// 入
		cnt1[s1[right]-'a']++

		left := right + 1 - len(s2)
		if left < 0 {
			// 窗口大小不够
			continue
		}

		// 更新
		if cnt1 == cnt2 {
			return true
		}

		// 出
		cnt1[s1[left]-'a']--
	}

	return false
}
