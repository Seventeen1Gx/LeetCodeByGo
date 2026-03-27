package LeetCode

func RepeatCharacter_1(s string) byte {
	hashSet := make(map[byte]bool)
	for _, ch := range []byte(s) {
		if hashSet[ch] {
			return ch
		}
		hashSet[ch] = true
	}
	return byte(0)
}

func RepeatCharacter_2(s string) byte {
	// 用一个整数(32位)来记录存在情况
	// 因为 ch 只在 26 个字母
	seen := int32(0)
	for _, ch := range []byte(s) {
		bit := ch - 'a'
		// 检查 seen 对应位是否为 1
		if seen>>bit&1 == 1 {
			return ch
		}
		// 将 seen 对应位置为 1
		seen |= 1 << int32(bit)
	}
	return byte(0)
}
