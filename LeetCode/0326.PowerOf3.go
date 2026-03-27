package LeetCode

import "math"

func IsPowerOfThree_1(x int) bool {
	if x <= 0 {
		return false
	}
	for x%3 == 0 {
		x /= 3
	}
	return x == 1
}

func IsPowerOfThree_2(x int) bool {
	// 找到最大的 3 的幂
	n := 3
	for {
		if n*3 > math.MaxInt32 {
			break
		}
		n *= 3
	}
	return x > 0 && n%x == 0
}
