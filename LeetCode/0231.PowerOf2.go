package LeetCode

func IsPowerOfTwo(n int) bool {
	// 10 100 1000 ..
	return n != 0 && n&(n-1) == 0
}
