package LeetCode

import "sort"

func SplitDigits(num int) int {
	digits := make([]int, 4)

	for i := 0; i < 4; i++ {
		digits[0] = num % 10
		num /= 10
	}

	sort.Ints(digits)

	// 让小数作高位
	// 1+3 的模式：digits[0]*100+digits[1]*10+digits[2]+digits[3]
	// 显然比下面的答案大
	return digits[0]*10 + digits[2] + digits[1]*10 + digits[3]
}
