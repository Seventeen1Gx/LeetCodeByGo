package LeetCode

func FirstUniqChar_1(s string) int {
	var count = make(map[byte]int)
	for _, ch := range []byte(s) {
		count[ch]++
	}
	for i, ch := range []byte(s) {
		if count[ch] == 1 {
			return i
		}
	}
	return -1
}

func FirstUniqChar_2(s string) int {
	// 使用队列，队列开头是候选的答案
	var queue []int
	var count = make(map[byte]int)
	for i, ch := range []byte(s) {
		if count[ch] == 0 {
			// ch 首次出现可能是答案，索引加入队列
			queue = append(queue, i)
			count[ch]++
		} else {
			count[ch]++
			// ch 非首次出现，队头元素有风险，检查并移出
			for len(queue) > 0 && count[s[queue[0]]] > 1 {
				queue = queue[1:]
			}
		}
	}

	if len(queue) > 0 {
		return queue[0]
	}
	return -1
}
